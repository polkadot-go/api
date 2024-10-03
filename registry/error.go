package registry

import libErr "github.com/polkadot-go/api/v4/error"

const (
	ErrRecursiveDecodersResolving            = libErr.Error("recursive decoders resolving")
	ErrErrorsTypeNotFound                    = libErr.Error("errors type not found")
	ErrErrorsTypeNotVariant                  = libErr.Error("errors type not a variant")
	ErrErrorFieldsRetrieval                  = libErr.Error("error fields retrieval")
	ErrCallsTypeNotFound                     = libErr.Error("calls type not found")
	ErrCallsTypeNotVariant                   = libErr.Error("calls type not a variant")
	ErrCallFieldsRetrieval                   = libErr.Error("call fields retrieval")
	ErrEventsTypeNotFound                    = libErr.Error("events type not found")
	ErrEventsTypeNotVariant                  = libErr.Error("events type not a variant")
	ErrEventFieldsRetrieval                  = libErr.Error("event fields retrieval")
	ErrFieldDecoderForRecursiveFieldNotFound = libErr.Error("field decoder for recursive field not found")
	ErrRecursiveFieldResolving               = libErr.Error("recursive field resolving")
	ErrFieldTypeNotFound                     = libErr.Error("field type not found")
	ErrFieldDecoderRetrieval                 = libErr.Error("field decoder retrieval")
	ErrCompactFieldTypeNotFound              = libErr.Error("compact field type not found")
	ErrCompositeTypeFieldsRetrieval          = libErr.Error("composite type fields retrieval")
	ErrArrayFieldTypeNotFound                = libErr.Error("array field type not found")
	ErrVectorFieldTypeNotFound               = libErr.Error("vector field type not found")
	ErrFieldTypeDefinitionNotSupported       = libErr.Error("field type definition not supported")
	ErrVariantTypeFieldsRetrieval            = libErr.Error("variant type fields decoding")
	ErrCompactTupleItemTypeNotFound          = libErr.Error("compact tuple item type not found")
	ErrCompactTupleItemFieldDecoderRetrieval = libErr.Error("compact tuple item field decoder retrieval")
	ErrCompactCompositeFieldTypeNotFound     = libErr.Error("compact composite field type not found")
	ErrCompactCompositeFieldDecoderRetrieval = libErr.Error("compact composite field decoder retrieval")
	ErrArrayItemFieldDecoderRetrieval        = libErr.Error("array item field decoder retrieval")
	ErrSliceItemFieldDecoderRetrieval        = libErr.Error("slice item field decoder retrieval")
	ErrTupleItemTypeNotFound                 = libErr.Error("tuple item type not found")
	ErrTupleItemFieldDecoderRetrieval        = libErr.Error("tuple item field decoder retrieval")
	ErrBitStoreTypeNotFound                  = libErr.Error("bit store type not found")
	ErrBitStoreTypeNotSupported              = libErr.Error("bit store type not supported")
	ErrBitOrderTypeNotFound                  = libErr.Error("bit order type not found")
	ErrBitOrderCreation                      = libErr.Error("bit order creation")
	ErrPrimitiveTypeNotSupported             = libErr.Error("primitive type not supported")
	ErrTypeFieldDecoding                     = libErr.Error("type field decoding")
	ErrVariantByteDecoding                   = libErr.Error("variant byte decoding")
	ErrVariantFieldDecoderNotFound           = libErr.Error("variant field decoder not found")
	ErrArrayItemDecoderNotFound              = libErr.Error("array item decoder not found")
	ErrArrayItemDecoding                     = libErr.Error("array item decoding")
	ErrSliceItemDecoderNotFound              = libErr.Error("slice item decoder not found")
	ErrSliceLengthDecoding                   = libErr.Error("slice length decoding")
	ErrSliceItemDecoding                     = libErr.Error("slice item decoding")
	ErrCompositeFieldDecoding                = libErr.Error("composite field decoding")
	ErrValueDecoding                         = libErr.Error("value decoding")
	ErrRecursiveFieldDecoderNotFound         = libErr.Error("recursive field decoder not found")
	ErrBitVecDecoding                        = libErr.Error("bit vec decoding")
	ErrNilTypeDecoder                        = libErr.Error("nil type decoder")
	ErrNilField                              = libErr.Error("nil field")
	ErrNilFieldDecoder                       = libErr.Error("nil field decoder")
	ErrDecodedFieldNotFound                  = libErr.Error("decoded field not found")
	ErrDecodedFieldValueTypeMismatch         = libErr.Error("decoded field value type mismatch")
	ErrDecodedFieldValueProcessingError      = libErr.Error("decoded field value processing error")
	ErrDecodedFieldValueNotAGenericSlice     = libErr.Error("decoded field value is not a generic slice")
	ErrExtrinsicFieldRetrieval               = libErr.Error("extrinsic field retrieval")
	ErrInvalidExtrinsicParams                = libErr.Error("invalid extrinsic params")
	ErrInvalidExtrinsicType                  = libErr.Error("invalid extrinsic type")
	ErrInvalidGenericExtrinsicType           = libErr.Error("invalid generic extrinsic type")
	ErrNilExtrinsicDecoder                   = libErr.Error("nil type decoder")
	ErrExtrinsicFieldNotFound                = libErr.Error("extrinsic field not found")
	ErrExtrinsicCompactLengthDecoding        = libErr.Error("extrinsic compact length decoding")
	ErrExtrinsicVersionDecoding              = libErr.Error("extrinsic version decoding")
	ErrUnexpectedExtrinsicParam              = libErr.Error("unexpected extrinsic param")
	ErrExtrinsicFieldDecoding                = libErr.Error("extrinsic field decoding")
)
