package constans

const (
	CODE   = "code"
	MSG    = "msg"
	RESULT = "result"
)

func OpSuccess() map[string]any {
	return map[string]any{
		CODE:   200,
		MSG:    "ok",
		RESULT: nil,
	}
}

func OpSuccessData(data any) map[string]any {
	return map[string]any{
		CODE:   200,
		MSG:    "ok",
		RESULT: data,
	}
}

func OpFail(msg string) map[string]any {
	return map[string]any{
		CODE:   -1,
		MSG:    msg,
		RESULT: nil,
	}
}
