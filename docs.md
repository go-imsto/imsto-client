# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/protobuf-spec/image.proto](#api/protobuf-spec/image.proto)
    - [FetchInput](#impb.FetchInput)
    - [ImageInput](#impb.ImageInput)
    - [ImageMeta](#impb.ImageMeta)
    - [ImageOutput](#impb.ImageOutput)
  
  
  
    - [ImageSvc](#impb.ImageSvc)
  

- [Scalar Value Types](#scalar-value-types)



<a name="api/protobuf-spec/image.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/protobuf-spec/image.proto



<a name="impb.FetchInput"></a>

### FetchInput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apiKey | [string](#string) |  | APIKey |
| uri | [string](#string) |  | full URI of image |
| referer | [string](#string) |  | special referer of request |
| roof | [string](#string) |  | section name |
| userID | [int64](#int64) |  | user ID |
| sizeOp | [string](#string) |  | size op |






<a name="impb.ImageInput"></a>

### ImageInput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apiKey | [string](#string) |  | APIKey |
| image | [bytes](#bytes) |  | Image Binary |
| name | [string](#string) |  | FileName |
| roof | [string](#string) |  | section name |
| userID | [int64](#int64) |  | user ID |
| sizeOp | [string](#string) |  | size op |






<a name="impb.ImageMeta"></a>

### ImageMeta



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| width | [int32](#int32) |  | 宽度 |
| height | [int32](#int32) |  | 高度 |
| quality | [int32](#int32) |  | 质量（JPEG） |
| size | [int32](#int32) |  | 大小 |
| ext | [string](#string) |  | 扩展名 |
| mime | [string](#string) |  | Mime type |






<a name="impb.ImageOutput"></a>

### ImageOutput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  | 路径 id.ext |
| uri | [string](#string) |  | URI 地址 |
| host | [string](#string) |  | 主机名，取自 stagHost |
| ID | [string](#string) |  | ID 编码 |
| meta | [ImageMeta](#impb.ImageMeta) |  | 图片 Meta |





 

 

 


<a name="impb.ImageSvc"></a>

### ImageSvc


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Fetch | [FetchInput](#impb.FetchInput) | [ImageOutput](#impb.ImageOutput) | Fetch |
| Store | [ImageInput](#impb.ImageInput) | [ImageOutput](#impb.ImageOutput) | Store |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

