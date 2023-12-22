## tsel
``
2023/12/20 16:30:05 &{POST https://secure.ximpay.com/api/dev04/Gopayment.aspx HTTP/1.1 1 1 map[Content-Type:[application/json]] {{"partnerid":"SEHAT","itemid":"SHT00001","cbparam":"123456","token":"286e21972445862cfdf927edbc90b698","op":"TSEL","msisdn":"081299708787"}} 0x73f2e0 139 [] false secure.ximpay.com map[] map[] <nil> map[]   <nil> <nil> <nil> {{}}}
2023/12/20 16:30:06 {"ximpaytransaction":[{"responsecode":1,"ximpayid":"D94478CAB24E47948832133FFAB55CB7","ximpayitem":[{"name":"Item 2K","price":"2.220"}]}]}
``

## h3i
``
2023/12/20 16:43:51 &{POST https://secure.ximpay.com/api/dev03/Gopayment.aspx HTTP/1.1 1 1 map[Content-Type:[application/json]] {{"partnerid":"SEHAT","itemid":"SHT00001","cbparam":"12","token":"5f3b8bc6866cdbde3a884817a6d4c470","op":"HTI","msisdn":"6289513144793"}} 0x73f2e0 135 [] false secure.ximpay.com map[] map[] <nil> map[]   <nil> <nil> <nil> {{}}}
2023/12/20 16:43:52 {"ximpaytransaction":[{"responsecode":1,"ximpayid":"BB06FB969F5F45439D33E193F25FD24C","ximpayitem":[{"name":"Item 2K","price":"Rp.2,200"}]}]}
``

## xl
``
2023/12/20 16:50:12 &{POST https://secure.ximpay.com/api/dev08flex/Gopayment.aspx HTTP/1.1 1 1 map[Content-Type:[application/json]] {{"partnerid":"SEHAT","item_name":"Item 2K","item_desc":"Item 2K","amount":2220,"cbparam":"123","token":"647a26c309829bd1a4df0bd988c18f6d","op":"xl","msisdn":"087810001000"}} 0x73f2e0 172 [] false secure.ximpay.com map[] map[] <nil> map[]   <nil> <nil> <nil> {{}}}
2023/12/20 16:50:12 {"ximpaytransaction":[{"responsecode":1,"ximpayid":"4E0840FAF83E44468BBEAF8518AA6605","ximpayitem":[{"item_name":"Item 2K","item_desc":"Item 2K","amount":2220}]}]}
``

## indosat
``
2023/12/20 16:52:59 &{POST https://secure.ximpay.com/api/dev07flex/Gopayment.aspx HTTP/1.1 1 1 map[Content-Type:[application/json]] {{"partnerid":"SEHAT","item_name":"Item 2K","item_desc":"Item 2K","amount":2220,"charge_type":"ISAT_GENERAL","cbparam":"1234","token":"3b59360289bf85394d298d22f552bda4","op":"ISAT","msisdn":"085610001000"}} 0x73f2e0 204 [] false secure.ximpay.com map[] map[] <nil> map[]   <nil> <nil> <nil> {{}}}
2023/12/20 16:52:59 {"ximpaytransaction":[{"responsecode":1,"ximpayid":"3DA2183B0CBA47F98F2DC8377A992A8A","ximpayitem":[{"item_name":"Item 2K","item_desc":"Item 2K","amount":2220}]}]}
``

## smartfren
``
2023/12/20 16:57:30 &{POST https://secure.ximpay.com/api/dev10SDPflex/Gopayment.aspx HTTP/1.1 1 1 map[Content-Type:[application/json]] {{"partnerid":"SEHAT","item_name":"Item 2K","item_desc":"Item 2K","amount_exc":2220,"cbparam":"12345","token":"3cfdd6b6bfd30af74ca9191466b08464","op":"SF","msisdn":"088110002000"}} 0x73f2e0 178 [] false secure.ximpay.com map[] map[] <nil> map[]   <nil> <nil> <nil> {{}}}
2023/12/20 16:57:31 {"ximpaytransaction":[{"responsecode":1,"ximpayid":"BEC5C7E4A8E2417CBBB821ECD5B28164","ximpayitem":[{"item_name":"Item 2K","item_desc":"Item 2K","amount_exc":2220}]}]}
``
