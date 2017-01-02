// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: posts.go
// DO NOT EDIT!

package common

import (
	"encoding/base64"
	fflib "github.com/pquerna/ffjson/fflib/v1"
	"reflect"
)

func (mj *BoardThread) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *BoardThread) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"id":`)
	fflib.FormatBits2(buf, uint64(mj.ID), 10, false)
	buf.WriteString(`,"time":`)
	fflib.FormatBits2(buf, uint64(mj.Time), 10, mj.Time < 0)
	buf.WriteByte(',')
	if len(mj.Name) != 0 {
		buf.WriteString(`"name":`)
		fflib.WriteJsonString(buf, string(mj.Name))
		buf.WriteByte(',')
	}
	if len(mj.Trip) != 0 {
		buf.WriteString(`"trip":`)
		fflib.WriteJsonString(buf, string(mj.Trip))
		buf.WriteByte(',')
	}
	if len(mj.Auth) != 0 {
		buf.WriteString(`"auth":`)
		fflib.WriteJsonString(buf, string(mj.Auth))
		buf.WriteByte(',')
	}
	if mj.Image != nil {
		if true {
			buf.WriteString(`"image":`)

			{

				err = mj.Image.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if mj.Locked != false {
		if mj.Locked {
			buf.WriteString(`"locked":true`)
		} else {
			buf.WriteString(`"locked":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Archived != false {
		if mj.Archived {
			buf.WriteString(`"archived":true`)
		} else {
			buf.WriteString(`"archived":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Sticky != false {
		if mj.Sticky {
			buf.WriteString(`"sticky":true`)
		} else {
			buf.WriteString(`"sticky":false`)
		}
		buf.WriteByte(',')
	}
	buf.WriteString(`"postCtr":`)
	fflib.FormatBits2(buf, uint64(mj.PostCtr), 10, false)
	buf.WriteString(`,"imageCtr":`)
	fflib.FormatBits2(buf, uint64(mj.ImageCtr), 10, false)
	buf.WriteString(`,"replyTime":`)
	fflib.FormatBits2(buf, uint64(mj.ReplyTime), 10, mj.ReplyTime < 0)
	buf.WriteString(`,"lastUpdated":`)
	fflib.FormatBits2(buf, uint64(mj.LastUpdated), 10, mj.LastUpdated < 0)
	buf.WriteString(`,"subject":`)
	fflib.WriteJsonString(buf, string(mj.Subject))
	buf.WriteString(`,"board":`)
	fflib.WriteJsonString(buf, string(mj.Board))
	buf.WriteByte('}')
	return nil
}

func (mj *Command) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Command) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"type":`)
	fflib.FormatBits2(buf, uint64(mj.Type), 10, false)
	buf.WriteString(`,"val":`)
	/* Interface types must use runtime reflection. type=interface {} kind=interface */
	err = buf.Encode(mj.Val)
	if err != nil {
		return err
	}
	buf.WriteByte('}')
	return nil
}

func (mj *DatabasePost) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *DatabasePost) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "lastUpdated":`)
	fflib.FormatBits2(buf, uint64(mj.LastUpdated), 10, mj.LastUpdated < 0)
	buf.WriteString(`,"IP":`)
	fflib.WriteJsonString(buf, string(mj.IP))
	buf.WriteString(`,"Password":`)
	if mj.Password != nil {
		buf.WriteString(`"`)
		{
			enc := base64.NewEncoder(base64.StdEncoding, buf)
			enc.Write(reflect.Indirect(reflect.ValueOf(mj.Password)).Bytes())
			enc.Close()
		}
		buf.WriteString(`"`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"Log":`)
	if mj.Log != nil {
		buf.WriteString(`[`)
		for i, v := range mj.Log {
			if i != 0 {
				buf.WriteString(`,`)
			}
			fflib.WriteJsonString(buf, string(v))
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"op":`)
	fflib.FormatBits2(buf, uint64(mj.OP), 10, false)
	buf.WriteString(`,"board":`)
	fflib.WriteJsonString(buf, string(mj.Board))
	buf.WriteByte(',')
	if mj.Editing != false {
		if mj.Editing {
			buf.WriteString(`"editing":true`)
		} else {
			buf.WriteString(`"editing":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Deleted != false {
		if mj.Deleted {
			buf.WriteString(`"deleted":true`)
		} else {
			buf.WriteString(`"deleted":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Banned != false {
		if mj.Banned {
			buf.WriteString(`"banned":true`)
		} else {
			buf.WriteString(`"banned":false`)
		}
		buf.WriteByte(',')
	}
	buf.WriteString(`"id":`)
	fflib.FormatBits2(buf, uint64(mj.ID), 10, false)
	buf.WriteString(`,"time":`)
	fflib.FormatBits2(buf, uint64(mj.Time), 10, mj.Time < 0)
	buf.WriteString(`,"body":`)
	fflib.WriteJsonString(buf, string(mj.Body))
	buf.WriteByte(',')
	if len(mj.Name) != 0 {
		buf.WriteString(`"name":`)
		fflib.WriteJsonString(buf, string(mj.Name))
		buf.WriteByte(',')
	}
	if len(mj.Trip) != 0 {
		buf.WriteString(`"trip":`)
		fflib.WriteJsonString(buf, string(mj.Trip))
		buf.WriteByte(',')
	}
	if len(mj.Auth) != 0 {
		buf.WriteString(`"auth":`)
		fflib.WriteJsonString(buf, string(mj.Auth))
		buf.WriteByte(',')
	}
	if mj.Image != nil {
		if true {
			buf.WriteString(`"image":`)

			{

				err = mj.Image.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if len(mj.Backlinks) != 0 {
		/* Falling back. type=common.LinkMap kind=map */
		buf.WriteString(`"backlinks":`)
		err = buf.Encode(mj.Backlinks)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	if len(mj.Links) != 0 {
		/* Falling back. type=common.LinkMap kind=map */
		buf.WriteString(`"links":`)
		err = buf.Encode(mj.Links)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	if len(mj.Commands) != 0 {
		buf.WriteString(`"commands":`)
		if mj.Commands != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Commands {
				if i != 0 {
					buf.WriteString(`,`)
				}

				{

					err = v.MarshalJSONBuf(buf)
					if err != nil {
						return err
					}

				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

func (mj *DatabaseThread) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *DatabaseThread) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"PostCtr":`)
	fflib.FormatBits2(buf, uint64(mj.PostCtr), 10, false)
	buf.WriteString(`,"ImageCtr":`)
	fflib.FormatBits2(buf, uint64(mj.ImageCtr), 10, false)
	buf.WriteString(`,"ID":`)
	fflib.FormatBits2(buf, uint64(mj.ID), 10, false)
	buf.WriteString(`,"ReplyTime":`)
	fflib.FormatBits2(buf, uint64(mj.ReplyTime), 10, mj.ReplyTime < 0)
	buf.WriteString(`,"Subject":`)
	fflib.WriteJsonString(buf, string(mj.Subject))
	buf.WriteString(`,"Board":`)
	fflib.WriteJsonString(buf, string(mj.Board))
	buf.WriteByte('}')
	return nil
}

func (mj *Link) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Link) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"op":`)
	fflib.FormatBits2(buf, uint64(mj.OP), 10, false)
	buf.WriteString(`,"board":`)
	fflib.WriteJsonString(buf, string(mj.Board))
	buf.WriteByte('}')
	return nil
}

func (mj *Post) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Post) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if mj.Editing != false {
		if mj.Editing {
			buf.WriteString(`"editing":true`)
		} else {
			buf.WriteString(`"editing":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Deleted != false {
		if mj.Deleted {
			buf.WriteString(`"deleted":true`)
		} else {
			buf.WriteString(`"deleted":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Banned != false {
		if mj.Banned {
			buf.WriteString(`"banned":true`)
		} else {
			buf.WriteString(`"banned":false`)
		}
		buf.WriteByte(',')
	}
	buf.WriteString(`"id":`)
	fflib.FormatBits2(buf, uint64(mj.ID), 10, false)
	buf.WriteString(`,"time":`)
	fflib.FormatBits2(buf, uint64(mj.Time), 10, mj.Time < 0)
	buf.WriteString(`,"body":`)
	fflib.WriteJsonString(buf, string(mj.Body))
	buf.WriteByte(',')
	if len(mj.Name) != 0 {
		buf.WriteString(`"name":`)
		fflib.WriteJsonString(buf, string(mj.Name))
		buf.WriteByte(',')
	}
	if len(mj.Trip) != 0 {
		buf.WriteString(`"trip":`)
		fflib.WriteJsonString(buf, string(mj.Trip))
		buf.WriteByte(',')
	}
	if len(mj.Auth) != 0 {
		buf.WriteString(`"auth":`)
		fflib.WriteJsonString(buf, string(mj.Auth))
		buf.WriteByte(',')
	}
	if mj.Image != nil {
		if true {
			buf.WriteString(`"image":`)

			{

				err = mj.Image.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if len(mj.Backlinks) != 0 {
		/* Falling back. type=common.LinkMap kind=map */
		buf.WriteString(`"backlinks":`)
		err = buf.Encode(mj.Backlinks)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	if len(mj.Links) != 0 {
		/* Falling back. type=common.LinkMap kind=map */
		buf.WriteString(`"links":`)
		err = buf.Encode(mj.Links)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	if len(mj.Commands) != 0 {
		buf.WriteString(`"commands":`)
		if mj.Commands != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Commands {
				if i != 0 {
					buf.WriteString(`,`)
				}

				{

					err = v.MarshalJSONBuf(buf)
					if err != nil {
						return err
					}

				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

func (mj *StandalonePost) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *StandalonePost) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "op":`)
	fflib.FormatBits2(buf, uint64(mj.OP), 10, false)
	buf.WriteString(`,"board":`)
	fflib.WriteJsonString(buf, string(mj.Board))
	buf.WriteByte(',')
	if mj.Editing != false {
		if mj.Editing {
			buf.WriteString(`"editing":true`)
		} else {
			buf.WriteString(`"editing":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Deleted != false {
		if mj.Deleted {
			buf.WriteString(`"deleted":true`)
		} else {
			buf.WriteString(`"deleted":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Banned != false {
		if mj.Banned {
			buf.WriteString(`"banned":true`)
		} else {
			buf.WriteString(`"banned":false`)
		}
		buf.WriteByte(',')
	}
	buf.WriteString(`"id":`)
	fflib.FormatBits2(buf, uint64(mj.ID), 10, false)
	buf.WriteString(`,"time":`)
	fflib.FormatBits2(buf, uint64(mj.Time), 10, mj.Time < 0)
	buf.WriteString(`,"body":`)
	fflib.WriteJsonString(buf, string(mj.Body))
	buf.WriteByte(',')
	if len(mj.Name) != 0 {
		buf.WriteString(`"name":`)
		fflib.WriteJsonString(buf, string(mj.Name))
		buf.WriteByte(',')
	}
	if len(mj.Trip) != 0 {
		buf.WriteString(`"trip":`)
		fflib.WriteJsonString(buf, string(mj.Trip))
		buf.WriteByte(',')
	}
	if len(mj.Auth) != 0 {
		buf.WriteString(`"auth":`)
		fflib.WriteJsonString(buf, string(mj.Auth))
		buf.WriteByte(',')
	}
	if mj.Image != nil {
		if true {
			buf.WriteString(`"image":`)

			{

				err = mj.Image.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if len(mj.Backlinks) != 0 {
		/* Falling back. type=common.LinkMap kind=map */
		buf.WriteString(`"backlinks":`)
		err = buf.Encode(mj.Backlinks)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	if len(mj.Links) != 0 {
		/* Falling back. type=common.LinkMap kind=map */
		buf.WriteString(`"links":`)
		err = buf.Encode(mj.Links)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	if len(mj.Commands) != 0 {
		buf.WriteString(`"commands":`)
		if mj.Commands != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Commands {
				if i != 0 {
					buf.WriteString(`,`)
				}

				{

					err = v.MarshalJSONBuf(buf)
					if err != nil {
						return err
					}

				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

func (mj *Thread) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Thread) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteByte('{')
	if mj.Abbrev != false {
		if mj.Abbrev {
			buf.WriteString(`"abbrev":true`)
		} else {
			buf.WriteString(`"abbrev":false`)
		}
		buf.WriteByte(',')
	}
	buf.WriteString(`"posts":`)
	if mj.Posts != nil {
		buf.WriteString(`[`)
		for i, v := range mj.Posts {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				err = v.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte(',')
	if mj.Editing != false {
		if mj.Editing {
			buf.WriteString(`"editing":true`)
		} else {
			buf.WriteString(`"editing":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Deleted != false {
		if mj.Deleted {
			buf.WriteString(`"deleted":true`)
		} else {
			buf.WriteString(`"deleted":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Banned != false {
		if mj.Banned {
			buf.WriteString(`"banned":true`)
		} else {
			buf.WriteString(`"banned":false`)
		}
		buf.WriteByte(',')
	}
	buf.WriteString(`"id":`)
	fflib.FormatBits2(buf, uint64(mj.ID), 10, false)
	buf.WriteString(`,"time":`)
	fflib.FormatBits2(buf, uint64(mj.Time), 10, mj.Time < 0)
	buf.WriteString(`,"body":`)
	fflib.WriteJsonString(buf, string(mj.Body))
	buf.WriteByte(',')
	if len(mj.Name) != 0 {
		buf.WriteString(`"name":`)
		fflib.WriteJsonString(buf, string(mj.Name))
		buf.WriteByte(',')
	}
	if len(mj.Trip) != 0 {
		buf.WriteString(`"trip":`)
		fflib.WriteJsonString(buf, string(mj.Trip))
		buf.WriteByte(',')
	}
	if len(mj.Auth) != 0 {
		buf.WriteString(`"auth":`)
		fflib.WriteJsonString(buf, string(mj.Auth))
		buf.WriteByte(',')
	}
	if mj.Image != nil {
		if true {
			buf.WriteString(`"image":`)

			{

				err = mj.Image.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if len(mj.Backlinks) != 0 {
		/* Falling back. type=common.LinkMap kind=map */
		buf.WriteString(`"backlinks":`)
		err = buf.Encode(mj.Backlinks)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	if len(mj.Links) != 0 {
		/* Falling back. type=common.LinkMap kind=map */
		buf.WriteString(`"links":`)
		err = buf.Encode(mj.Links)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	if len(mj.Commands) != 0 {
		buf.WriteString(`"commands":`)
		if mj.Commands != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Commands {
				if i != 0 {
					buf.WriteString(`,`)
				}

				{

					err = v.MarshalJSONBuf(buf)
					if err != nil {
						return err
					}

				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.Locked != false {
		if mj.Locked {
			buf.WriteString(`"locked":true`)
		} else {
			buf.WriteString(`"locked":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Archived != false {
		if mj.Archived {
			buf.WriteString(`"archived":true`)
		} else {
			buf.WriteString(`"archived":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Sticky != false {
		if mj.Sticky {
			buf.WriteString(`"sticky":true`)
		} else {
			buf.WriteString(`"sticky":false`)
		}
		buf.WriteByte(',')
	}
	buf.WriteString(`"postCtr":`)
	fflib.FormatBits2(buf, uint64(mj.PostCtr), 10, false)
	buf.WriteString(`,"imageCtr":`)
	fflib.FormatBits2(buf, uint64(mj.ImageCtr), 10, false)
	buf.WriteString(`,"replyTime":`)
	fflib.FormatBits2(buf, uint64(mj.ReplyTime), 10, mj.ReplyTime < 0)
	buf.WriteString(`,"lastUpdated":`)
	fflib.FormatBits2(buf, uint64(mj.LastUpdated), 10, mj.LastUpdated < 0)
	buf.WriteString(`,"subject":`)
	fflib.WriteJsonString(buf, string(mj.Subject))
	buf.WriteString(`,"board":`)
	fflib.WriteJsonString(buf, string(mj.Board))
	buf.WriteByte('}')
	return nil
}

func (mj *ThreadCommon) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *ThreadCommon) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteByte('{')
	if mj.Locked != false {
		if mj.Locked {
			buf.WriteString(`"locked":true`)
		} else {
			buf.WriteString(`"locked":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Archived != false {
		if mj.Archived {
			buf.WriteString(`"archived":true`)
		} else {
			buf.WriteString(`"archived":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Sticky != false {
		if mj.Sticky {
			buf.WriteString(`"sticky":true`)
		} else {
			buf.WriteString(`"sticky":false`)
		}
		buf.WriteByte(',')
	}
	buf.WriteString(`"postCtr":`)
	fflib.FormatBits2(buf, uint64(mj.PostCtr), 10, false)
	buf.WriteString(`,"imageCtr":`)
	fflib.FormatBits2(buf, uint64(mj.ImageCtr), 10, false)
	buf.WriteString(`,"replyTime":`)
	fflib.FormatBits2(buf, uint64(mj.ReplyTime), 10, mj.ReplyTime < 0)
	buf.WriteString(`,"lastUpdated":`)
	fflib.FormatBits2(buf, uint64(mj.LastUpdated), 10, mj.LastUpdated < 0)
	buf.WriteString(`,"subject":`)
	fflib.WriteJsonString(buf, string(mj.Subject))
	buf.WriteString(`,"board":`)
	fflib.WriteJsonString(buf, string(mj.Board))
	buf.WriteByte('}')
	return nil
}
