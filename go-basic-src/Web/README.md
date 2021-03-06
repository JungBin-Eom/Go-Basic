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
  * REST: Representational State Transfer
  * URL(URI)에 method(GET, POST, PUT, DELETE...)와 함께 표시하여 CRUD 구현
- 같은 경로로 요청을 보내도 method에 따라 핸들러 달라져야함(당연히 하는 동작이 다르기 때문)
- update시 원하는 필드만 업데이트 할 수 있도록 struct를 새로 만듬
### Web4(Decorator)
- Decorator란?
  * 프로그래밍 패턴 중 하나
  * 꾸미기 역할(기본 기능은 갖추고 있으면서 부가기능을 추가)
  * 잘 바뀌는 부가기능의 특성 상 기능별로 따로 분리하여 만들어야 함
    + 하나로 묶어서 만들면 부가기능이 수정될 때마다 모든 기능을 바꾸어야 함
- 하나의 인터페이스를 구현한 구조체를 내장하여 전달
- 웹서버에서 decorator 패턴이 필요한 이유
  * 문서를 만들어서 response를 전달하는 기본기능에 부가기능을 더하여 보냄
    + 암호화, 압축, log 기록, 분석 자료 전송 등을 부가기능이라고 함
### Web5(Template)
- html/template 패키지 사용
- HTML을 만들 때 내용을 채움
  * 변경되지 않는 부분은 template화
  * 변경되는 부분을 바꾸어 채우기
- html/template은 특수기호 탈락
  * <script>태그 안에서는 특수기호 유지
- text/template은 특수기호 유지
### Web6(Packages)
- gorilla/pat: 메서드에 따른 핸들러 추가 간단하게 해줌
- unrolled/render: JSON/XML/binary/HTML 템플릿 렌더링 쉽게 해줌
- urfave/negroni: HTTP 미들웨어로 많이 쓰이는 부가 기능 추가 쉽게 해줌
### Web7(EventSource)
- WebSocket은 TCP 채널을 생성하여 연결을 끊지 않고 send/receive를 가능하게 해주는 통신 프로토콜
- EventSource는 연결을 끊지 않고 서버에서 이벤트를 보내도록 해주는 인터페이스
  * push 알림 등에 사용
### Web8(OAuth 2.0)
- 앱이나 사이트에서 자체 회원가입 시스템 없이 외부 소셜 사이트(goole, facebook, kakao 등)로부터 회원 정보를 가져와 가입을 제공
- 개인정보(ID, 비밀번호 등) 보호가 어려워짐으로써 큰 회사의 정보를 사용
- OAuth 요청시 API Key를 제공하고 이 key를 통해 회원 정보에 접근 가능
- Callback 함수를 통해 로그인 제공
### Web9(배포)
- Domain 획득 -> DNS 등록 -> IP주소와 Domain 매핑
- 가상화 기술의 발전으로 클라우드 서비스 발달
- 인프라 자원을 제공(IaaS)
  * 가상 머신(자원)을 빌려서 웹서버 구동
- 웹 서버를 구동할 수 있는 플랫폼 제공(PaaS)
  * 인프라는 신경쓰지 않아도 됨
  * 디테일한 설정 불가능
- 서비스 형태로 접근할 수 있도록 제공(SaaS)