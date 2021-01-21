# Go 언어로 만드는 웹
`본 디렉토리의 소스코드는 Youtube 채널 'Tucker Programming'님의 강의를 참고하여 작성하였습니다.`  
  
---

### Web1(Basic)
- HTML: Hyper Text Markup Language
- HTTP: Hyper Text Transfer Protocol
  * Hyper Text를 주고받는 통신 규약
  * request/response를 어떻게 받을 것이냐? 하는 규약
- 요즘은 Client Rendering으로 웹페이지 구성
- Server에서 HTML로 전체 페이지를 주는 것이 아니라 부분 혹은 데이터(JSON) 전달
- HandleFunc: function 형태로 직접 등록
- Handle: Handler 인터페이스를 구현한 인스턴스를 등록
- Mux로 핸들러 인스턴스를 등록하여 동적으로 전달
- URL에 있는 데이터를 쿼리로 얻을 수 있음
- JSON을 통해서 데이터를 주고받음
  * Decode(JSON -> Object) - Decode나 Unmarshal 사용
  * Encode(Object -> JSON) - Encode나 Marshal 사용
  * 헤더에 content-name 명시
- Test 코드 작성하여 테스트하는 습관 기륵
  * goconvey 실행하여 웹 서버를 구동하고 테스트 수행
    + 일일이 브라우저 실행해보지 않고 백그라운드에서 테스트수행
  * assert 패키지로 쉽게 테스트 코드 작성
### Web2(File)
- http.FileServer(http.Dir("<directory_name>"))으로 파일서버 구축
- 지정한 폴더 내에 웹페이지를 만듦
- 웹 페이지에서의 동작을 통해 서버로 request
  * action을 수행할 때 어떤 경로로 갈 것인지 결정
  * 핸들러 작성하여 해당 경로에 등록
- 생성(오픈)한 파일은 defer로 프로그램 종료 전 항상 닫아줄 것
- 업로드 한 파일과 서버로 전송된 파일이 같은지 byte array를 비교하는 과정도 필요
### Web3(RESTful API)
- RESTful API란?
  + REST: Representational State Transfer
  + URL(URI)에 method(GET, POST, PUT, DELETE...)와 함께 표시하여 CRUD 구현
- 같은 경로로 요청을 보내도 method에 따라 핸들러 달라져야함(당연히 하는 동작이 다르기 때문)
- update시 원하는 필드만 업데이트 할 수 있도록 struct를 새로 만듬