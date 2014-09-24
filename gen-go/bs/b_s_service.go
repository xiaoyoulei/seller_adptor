// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package bs

import (
	"fmt"
//	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/apache/thrift/lib/go/thrift"
	"math"
)

// (needed to ensure safety because of naive import list construction.)
var _ = math.MinInt32
var _ = thrift.ZERO
var _ = fmt.Printf

type BSService interface {
	Search2() (err error)
	// Parameters:
	//  - Req
	Search(req *BSRequest) (r *BSResponse, err error)
}

type BSServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewBSServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *BSServiceClient {
	return &BSServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewBSServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *BSServiceClient {
	return &BSServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

func (p *BSServiceClient) Search2() (err error) {
	if err = p.sendSearch2(); err != nil {
		return
	}
	return p.recvSearch2()
}

func (p *BSServiceClient) sendSearch2() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("search2", thrift.CALL, p.SeqId)
	args1 := NewSearch2Args()
	err = args1.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *BSServiceClient) recvSearch2() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error3 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error4 error
		error4, err = error3.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error4
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result2 := NewSearch2Result()
	err = result2.Read(iprot)
	iprot.ReadMessageEnd()
	return
}

// Parameters:
//  - Req
func (p *BSServiceClient) Search(req *BSRequest) (r *BSResponse, err error) {
	if err = p.sendSearch(req); err != nil {
		return
	}
	return p.recvSearch()
}

func (p *BSServiceClient) sendSearch(req *BSRequest) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("search", thrift.CALL, p.SeqId)
	args5 := NewSearchArgs()
	args5.Req = req
	err = args5.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *BSServiceClient) recvSearch() (value *BSResponse, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error7 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error8 error
		error8, err = error7.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error8
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result6 := NewSearchResult()
	err = result6.Read(iprot)
	iprot.ReadMessageEnd()
	value = result6.Success
	return
}

type BSServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      BSService
}

func (p *BSServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *BSServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *BSServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewBSServiceProcessor(handler BSService) *BSServiceProcessor {

	self9 := &BSServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self9.processorMap["search2"] = &bSServiceProcessorSearch2{handler: handler}
	self9.processorMap["search"] = &bSServiceProcessorSearch{handler: handler}
	return self9
}

func (p *BSServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x10 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x10.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x10

}

type bSServiceProcessorSearch2 struct {
	handler BSService
}

func (p *bSServiceProcessorSearch2) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewSearch2Args()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("search2", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewSearch2Result()
	if err = p.handler.Search2(); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing search2: "+err.Error())
		oprot.WriteMessageBegin("search2", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("search2", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type bSServiceProcessorSearch struct {
	handler BSService
}

func (p *bSServiceProcessorSearch) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewSearchArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("search", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewSearchResult()
	if result.Success, err = p.handler.Search(args.Req); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing search: "+err.Error())
		oprot.WriteMessageBegin("search", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("search", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type Search2Args struct {
}

func NewSearch2Args() *Search2Args {
	return &Search2Args{}
}

func (p *Search2Args) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Search2Args) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("search2_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *Search2Args) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Search2Args(%+v)", *p)
}

type Search2Result struct {
}

func NewSearch2Result() *Search2Result {
	return &Search2Result{}
}

func (p *Search2Result) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Search2Result) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("search2_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *Search2Result) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Search2Result(%+v)", *p)
}

type SearchArgs struct {
	Req *BSRequest `thrift:"req,1"`
}

func NewSearchArgs() *SearchArgs {
	return &SearchArgs{}
}

func (p *SearchArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *SearchArgs) readField1(iprot thrift.TProtocol) error {
	p.Req = NewBSRequest()
	if err := p.Req.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Req)
	}
	return nil
}

func (p *SearchArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("search_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *SearchArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Req != nil {
		if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:req: %s", p, err)
		}
		if err := p.Req.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Req)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:req: %s", p, err)
		}
	}
	return err
}

func (p *SearchArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SearchArgs(%+v)", *p)
}

type SearchResult struct {
	Success *BSResponse `thrift:"success,0"`
}

func NewSearchResult() *SearchResult {
	return &SearchResult{}
}

func (p *SearchResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *SearchResult) readField0(iprot thrift.TProtocol) error {
	p.Success = NewBSResponse()
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success)
	}
	return nil
}

func (p *SearchResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("search_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *SearchResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.Success != nil {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *SearchResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SearchResult(%+v)", *p)
}
