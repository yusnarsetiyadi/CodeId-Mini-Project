package middleware

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmhttp/v2"
	"go.elastic.co/apm/module/apmlogrus/v2"
	"go.elastic.co/apm/v2"

	"compass_mini_api/internal/abstraction"
)

// TraceJSON ...
type TraceJSON struct {
	// ClientIP equals Context's ClientIP method.
	ClientIP string `json:"client_ip"`
	// Session ...
	Session interface{} `json:"session"`
	// Request body
	Request interface{} `json:"request"`
	// Response body
	Response interface{} `json:"response"`
}

func init() {
	logrus.AddHook(&apmlogrus.Hook{
		LogLevels: logrus.AllLevels,
	})
}

// Trace ...
func Trace(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := c.Request()

		tracer := apm.DefaultTracer()
		reqIgnore := apmhttp.NewDynamicServerRequestIgnorer(tracer)
		if !tracer.Recording() || reqIgnore(req) {
			return next(c)
		}

		// Path
		path := req.URL.Path
		raw := req.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		tx, body, req := apmhttp.StartTransactionWithBody(tracer, req.Method+" "+path, req)
		defer tx.End()

		c.SetRequest(req)
		resp := c.Response()

		var (
			reqBody []byte
			resBody = new(bytes.Buffer)
		)
		defer func() {
			if v := recover(); v != nil {
				var ok bool
				if err, ok = v.(error); !ok {
					err = errors.New(fmt.Sprint(v))
				}
				c.Error(err)

				e := tracer.Recovered(v)
				e.SetTransaction(tx)
				setContext(&e.Context, req, resp, body)
				e.Send()
			}
			if err != nil {
				e := tracer.NewError(err)
				setContext(&e.Context, req, resp, body)
				e.SetTransaction(tx)
				e.Handled = true
				e.Send()
			}
			tx.Result = apmhttp.StatusCodeResult(resp.Status)
			if tx.Sampled() {
				if ctx, ok := c.(*abstraction.Context); ok && ctx != nil && ctx.Auth != nil {
					tx.Context.SetUserID(strconv.Itoa(ctx.Auth.ID))
				}
				tx.Context.SetCustom("request", fmt.Sprint(parseBody(reqBody)))
				tx.Context.SetCustom("response", fmt.Sprint(parseBody(resBody.Bytes())))
				setContext(&tx.Context, req, resp, body)
			}
			body.Discard()
		}()

		// Request
		if req.Body != nil { // Read
			reqBody, _ = io.ReadAll(req.Body)
		}
		req.Body = io.NopCloser(bytes.NewBuffer(reqBody)) // Reset

		// Response
		mw := io.MultiWriter(c.Response().Writer, resBody)
		writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: resp.Writer}
		resp.Writer = writer

		err = next(c)
		return
	}
}

func setContext(ctx *apm.Context, req *http.Request, resp *echo.Response, body *apm.BodyCapturer) {
	ctx.SetFramework("echo", echo.Version)
	ctx.SetHTTPRequest(req)
	ctx.SetHTTPRequestBody(body)
	ctx.SetHTTPStatusCode(resp.Status)
	ctx.SetHTTPResponseHeaders(resp.Header())
}

func parseBody(body []byte) (result interface{}) {
	_ = json.Unmarshal(body, &result)
	if res, ok := result.(map[string]interface{}); ok {
		res = hide(res)
		if data, ok := res["data"].(map[string]interface{}); ok {
			res["data"] = hide(data)
		}
		if data, ok := res["data"].([]interface{}); ok {
			res["data"] = fmt.Sprintf("[%d]", len(data))
		}
		b, _ := json.MarshalIndent(res, "", " ")
		result = string(b)
	}
	return
}

func hide(res map[string]interface{}) map[string]interface{} {
	if res["password"] != nil {
		res["password"] = "********"
	}
	if res["pin"] != nil {
		res["pin"] = "********"
	}
	if res["token"] != nil {
		res["token"] = "********"
	}
	if res["access_token"] != nil {
		res["access_token"] = "********"
	}
	if res["refresh_token"] != nil {
		res["refresh_token"] = "********"
	}
	return res
}

type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *bodyDumpResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *bodyDumpResponseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *bodyDumpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}
