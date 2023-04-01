// Code generated by thriftgo (0.2.8). DO NOT EDIT.

package goods

import (
	"NihiStore/server/cmd/api/biz/model/base"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type CreateGoodsRequest struct {
	Username         string      `thrift:"username,1" form:"username" json:"username" query:"username"`
	GoodsInformation *base.Goods `thrift:"goods_information,2" form:"goods_information" json:"goods_information" query:"goods_information"`
}

func NewCreateGoodsRequest() *CreateGoodsRequest {
	return &CreateGoodsRequest{}
}

func (p *CreateGoodsRequest) GetUsername() (v string) {
	return p.Username
}

var CreateGoodsRequest_GoodsInformation_DEFAULT *base.Goods

func (p *CreateGoodsRequest) GetGoodsInformation() (v *base.Goods) {
	if !p.IsSetGoodsInformation() {
		return CreateGoodsRequest_GoodsInformation_DEFAULT
	}
	return p.GoodsInformation
}

var fieldIDToName_CreateGoodsRequest = map[int16]string{
	1: "username",
	2: "goods_information",
}

func (p *CreateGoodsRequest) IsSetGoodsInformation() bool {
	return p.GoodsInformation != nil
}

func (p *CreateGoodsRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CreateGoodsRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *CreateGoodsRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Username = v
	}
	return nil
}

func (p *CreateGoodsRequest) ReadField2(iprot thrift.TProtocol) error {
	p.GoodsInformation = base.NewGoods()
	if err := p.GoodsInformation.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *CreateGoodsRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("CreateGoodsRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *CreateGoodsRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("username", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Username); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *CreateGoodsRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("goods_information", thrift.STRUCT, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.GoodsInformation.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *CreateGoodsRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreateGoodsRequest(%+v)", *p)
}

type CreateGoodsResponse struct {
	BaseResp *base.BaseResponse `thrift:"base_resp,1" form:"base_resp" json:"base_resp" query:"base_resp"`
}

func NewCreateGoodsResponse() *CreateGoodsResponse {
	return &CreateGoodsResponse{}
}

var CreateGoodsResponse_BaseResp_DEFAULT *base.BaseResponse

func (p *CreateGoodsResponse) GetBaseResp() (v *base.BaseResponse) {
	if !p.IsSetBaseResp() {
		return CreateGoodsResponse_BaseResp_DEFAULT
	}
	return p.BaseResp
}

var fieldIDToName_CreateGoodsResponse = map[int16]string{
	1: "base_resp",
}

func (p *CreateGoodsResponse) IsSetBaseResp() bool {
	return p.BaseResp != nil
}

func (p *CreateGoodsResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CreateGoodsResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *CreateGoodsResponse) ReadField1(iprot thrift.TProtocol) error {
	p.BaseResp = base.NewBaseResponse()
	if err := p.BaseResp.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *CreateGoodsResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("CreateGoodsResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *CreateGoodsResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("base_resp", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.BaseResp.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *CreateGoodsResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreateGoodsResponse(%+v)", *p)
}

type GoodsService interface {
	CreateGoods(ctx context.Context, req *CreateGoodsRequest) (r *CreateGoodsResponse, err error)
}

type GoodsServiceClient struct {
	c thrift.TClient
}

func NewGoodsServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *GoodsServiceClient {
	return &GoodsServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewGoodsServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *GoodsServiceClient {
	return &GoodsServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewGoodsServiceClient(c thrift.TClient) *GoodsServiceClient {
	return &GoodsServiceClient{
		c: c,
	}
}

func (p *GoodsServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *GoodsServiceClient) CreateGoods(ctx context.Context, req *CreateGoodsRequest) (r *CreateGoodsResponse, err error) {
	var _args GoodsServiceCreateGoodsArgs
	_args.Req = req
	var _result GoodsServiceCreateGoodsResult
	if err = p.Client_().Call(ctx, "CreateGoods", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type GoodsServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      GoodsService
}

func (p *GoodsServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *GoodsServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *GoodsServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewGoodsServiceProcessor(handler GoodsService) *GoodsServiceProcessor {
	self := &GoodsServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("CreateGoods", &goodsServiceProcessorCreateGoods{handler: handler})
	return self
}
func (p *GoodsServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type goodsServiceProcessorCreateGoods struct {
	handler GoodsService
}

func (p *goodsServiceProcessorCreateGoods) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GoodsServiceCreateGoodsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("CreateGoods", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := GoodsServiceCreateGoodsResult{}
	var retval *CreateGoodsResponse
	if retval, err2 = p.handler.CreateGoods(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing CreateGoods: "+err2.Error())
		oprot.WriteMessageBegin("CreateGoods", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("CreateGoods", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type GoodsServiceCreateGoodsArgs struct {
	Req *CreateGoodsRequest `thrift:"req,1"`
}

func NewGoodsServiceCreateGoodsArgs() *GoodsServiceCreateGoodsArgs {
	return &GoodsServiceCreateGoodsArgs{}
}

var GoodsServiceCreateGoodsArgs_Req_DEFAULT *CreateGoodsRequest

func (p *GoodsServiceCreateGoodsArgs) GetReq() (v *CreateGoodsRequest) {
	if !p.IsSetReq() {
		return GoodsServiceCreateGoodsArgs_Req_DEFAULT
	}
	return p.Req
}

var fieldIDToName_GoodsServiceCreateGoodsArgs = map[int16]string{
	1: "req",
}

func (p *GoodsServiceCreateGoodsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GoodsServiceCreateGoodsArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GoodsServiceCreateGoodsArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GoodsServiceCreateGoodsArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewCreateGoodsRequest()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *GoodsServiceCreateGoodsArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("CreateGoods_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GoodsServiceCreateGoodsArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GoodsServiceCreateGoodsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GoodsServiceCreateGoodsArgs(%+v)", *p)
}

type GoodsServiceCreateGoodsResult struct {
	Success *CreateGoodsResponse `thrift:"success,0,optional"`
}

func NewGoodsServiceCreateGoodsResult() *GoodsServiceCreateGoodsResult {
	return &GoodsServiceCreateGoodsResult{}
}

var GoodsServiceCreateGoodsResult_Success_DEFAULT *CreateGoodsResponse

func (p *GoodsServiceCreateGoodsResult) GetSuccess() (v *CreateGoodsResponse) {
	if !p.IsSetSuccess() {
		return GoodsServiceCreateGoodsResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_GoodsServiceCreateGoodsResult = map[int16]string{
	0: "success",
}

func (p *GoodsServiceCreateGoodsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GoodsServiceCreateGoodsResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GoodsServiceCreateGoodsResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GoodsServiceCreateGoodsResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewCreateGoodsResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *GoodsServiceCreateGoodsResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("CreateGoods_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GoodsServiceCreateGoodsResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *GoodsServiceCreateGoodsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GoodsServiceCreateGoodsResult(%+v)", *p)
}
