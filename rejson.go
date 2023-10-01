package rejson

import (
	"context"
	"github.com/nitishm/go-rejson/v4/rjs"
)

type Handler struct {
	clientName     string
	implementation ReJSON
}

func NewReJSONHandler() *Handler {
	return &Handler{clientName: rjs.ClientInactive}
}

// ReJSON provides an interface for various Go Redis Clients to implement ReJSON commands
type ReJSON interface {
	JSONSet(ctx context.Context, key, path string, obj interface{}, opts ...rjs.SetOption) (res interface{}, err error)

	JSONGet(ctx context.Context, key, path string, opts ...rjs.GetOption) (res interface{}, err error)

	JSONMGet(ctx context.Context, path string, keys ...string) (res interface{}, err error)

	JSONDel(ctx context.Context, key, path string) (res interface{}, err error)

	JSONType(ctx context.Context, key, path string) (res interface{}, err error)

	JSONNumIncrBy(ctx context.Context, key, path string, number int) (res interface{}, err error)

	JSONNumMultBy(ctx context.Context, key, path string, number int) (res interface{}, err error)

	JSONStrAppend(ctx context.Context, key, path string, jsonstring string) (res interface{}, err error)

	JSONStrLen(ctx context.Context, key, path string) (res interface{}, err error)

	JSONArrAppend(ctx context.Context, key, path string, values ...interface{}) (res interface{}, err error)

	JSONArrLen(ctx context.Context, key, path string) (res interface{}, err error)

	JSONArrPop(ctx context.Context, key, path string, index int) (res interface{}, err error)

	JSONArrIndex(ctx context.Context, key, path string, jsonValue interface{}, optionalRange ...int) (res interface{}, err error)

	JSONArrTrim(ctx context.Context, key, path string, start, end int) (res interface{}, err error)

	JSONArrInsert(ctx context.Context, key, path string, index int, values ...interface{}) (res interface{}, err error)

	JSONObjKeys(ctx context.Context, key, path string) (res interface{}, err error)

	JSONObjLen(ctx context.Context, key, path string) (res interface{}, err error)

	JSONDebug(ctx context.Context, subCmd rjs.DebugSubCommand, key, path string) (res interface{}, err error)

	JSONForget(ctx context.Context, key, path string) (res interface{}, err error)

	JSONResp(ctx context.Context, key, path string) (res interface{}, err error)
}

// JSONSet used to set a json object
//
// ReJSON syntax:
//
//	JSON.SET <key> <path> <json>
//			 [NX | XX]
func (r *Handler) JSONSet(ctx context.Context, key string, path string, obj interface{}, opts ...rjs.SetOption) (
	res interface{}, err error,
) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONSet(ctx, key, path, obj, opts...)
}

// JSONGet used to get a json object
//
// ReJSON syntax:
//
//	JSON.GET <key>
//			[INDENT indentation-string]
//			[NEWLINE line-break-string]
//			[SPACE space-string]
//			[NOESCAPE]
//			[path ...]
func (r *Handler) JSONGet(ctx context.Context, key, path string, opts ...rjs.GetOption) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONGet(ctx, key, path, opts...)
}

// JSONMGet used to get path values from multiple keys
//
// ReJSON syntax:
//
//	JSON.MGET <key> [key ...] <path>
func (r *Handler) JSONMGet(ctx context.Context, path string, keys ...string) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONMGet(ctx, path, keys...)
}

// JSONDel to delete a json object
//
// ReJSON syntax:
//
//	JSON.DEL <key> <path>
func (r *Handler) JSONDel(ctx context.Context, key string, path string) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONDel(ctx, key, path)
}

// JSONType to get the type of key or member at path.
//
// ReJSON syntax:
//
//	JSON.TYPE <key> [path]
func (r *Handler) JSONType(ctx context.Context, key, path string) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONType(ctx, key, path)
}

// JSONNumIncrBy to increment a number by provided amount
//
// ReJSON syntax:
//
//	JSON.NUMINCRBY <key> <path> <number>
func (r *Handler) JSONNumIncrBy(ctx context.Context, key, path string, number int) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONNumIncrBy(ctx, key, path, number)
}

// JSONNumMultBy to increment a number by provided amount
//
// ReJSON syntax:
//
//	JSON.NUMMULTBY <key> <path> <number>
func (r *Handler) JSONNumMultBy(ctx context.Context, key, path string, number int) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONNumMultBy(ctx, key, path, number)
}

// JSONStrAppend to append a jsonstring to an existing member
//
// ReJSON syntax:
//
//	JSON.STRAPPEND <key> [path] <json-string>
func (r *Handler) JSONStrAppend(ctx context.Context, key, path, jsonstring string) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONStrAppend(ctx, key, path, jsonstring)
}

// JSONStrLen to return the length of a string member
//
// ReJSON syntax:
//
//	JSON.STRLEN <key> [path]
func (r *Handler) JSONStrLen(ctx context.Context, key, path string) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONStrLen(ctx, key, path)
}

// JSONArrAppend to append json value into array at path
//
// ReJSON syntax:
//
//	JSON.ARRAPPEND <key> <path> <json> [json ...]
func (r *Handler) JSONArrAppend(ctx context.Context, key, path string, values ...interface{}) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONArrAppend(ctx, key, path, values...)
}

// JSONArrLen returns the length of the json array at path
//
// ReJSON syntax:
//
//	JSON.ARRLEN <key> [path]
func (r *Handler) JSONArrLen(ctx context.Context, key, path string) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONArrLen(ctx, key, path)
}

// JSONArrPop removes and returns element from the index in the array
// to pop last element use rejson.PopArrLast
//
// ReJSON syntax:
//
//	JSON.ARRPOP <key> [path [index]]
func (r *Handler) JSONArrPop(ctx context.Context, key, path string, index int) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONArrPop(ctx, key, path, index)
}

// JSONArrIndex returns the index of the json element provided and return -1 if element is not present
//
// ReJSON syntax:
//
//	JSON.ARRINDEX <key> <path> <json-scalar> [start [stop]]
func (r *Handler) JSONArrIndex(ctx context.Context, key, path string, jsonValue interface{}, optionalRange ...int) (
	res interface{}, err error,
) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONArrIndex(ctx, key, path, jsonValue, optionalRange...)
}

// JSONArrTrim trims an array so that it contains only the specified inclusive range of elements
//
// ReJSON syntax:
//
//	JSON.ARRTRIM <key> <path> <start> <stop>
func (r *Handler) JSONArrTrim(ctx context.Context, key, path string, start, end int) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONArrTrim(ctx, key, path, start, end)
}

// JSONArrInsert inserts the json value(s) into the array at path before the index (shifts to the right).
//
// ReJSON syntax:
//
//	JSON.ARRINSERT <key> <path> <index> <json> [json ...]
func (r *Handler) JSONArrInsert(ctx context.Context, key, path string, index int, values ...interface{}) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONArrInsert(ctx, key, path, index, values...)
}

// JSONObjKeys returns the keys in the object that's referenced by path
//
// ReJSON syntax:
//
//	JSON.OBJKEYS <key> [path]
func (r *Handler) JSONObjKeys(ctx context.Context, key, path string) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONObjKeys(ctx, key, path)
}

// JSONObjLen report the number of keys in the JSON Object at path in key
//
// ReJSON syntax:
//
//	JSON.OBJLEN <key> [path]
func (r *Handler) JSONObjLen(ctx context.Context, key, path string) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONObjLen(ctx, key, path)
}

// JSONDebug reports information
//
// ReJSON syntax:
//
//	JSON.DEBUG <subcommand & arguments>
//		JSON.DEBUG MEMORY <key> [path]	- report the memory usage in bytes of a value. path defaults to root if not provided.
//		JSON.DEBUG HELP					- reply with a helpful message
func (r *Handler) JSONDebug(ctx context.Context, subCmd rjs.DebugSubCommand, key, path string) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONDebug(ctx, subCmd, key, path)
}

// JSONForget is an alias for JSONDel
//
// ReJSON syntax:
//
//	JSON.FORGET <key> [path]
func (r *Handler) JSONForget(ctx context.Context, key, path string) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONForget(ctx, key, path)
}

// JSONResp returns the JSON in key in Redis Serialization Protocol (RESP).
//
// ReJSON syntax:
//
//	JSON.RESP <key> [path]
func (r *Handler) JSONResp(ctx context.Context, key, path string) (res interface{}, err error) {
	if r.clientName == rjs.ClientInactive {
		return nil, rjs.ErrNoClientSet
	}
	return r.implementation.JSONResp(ctx, key, path)
}
