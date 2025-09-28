---
title: ai-chat
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.30"

---

# ai-chat

Base URLs:

# Authentication

- HTTP Authentication, scheme: bearer

# character

## GET 角色列表

GET /api/v1/characters

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "message": "OK",
  "data": {
    "list": [
      {
        "id": 1,
        "name": "孙悟空",
        "subtitle": "齐天大圣",
        "description": "勇敢无畏、神通广大。与你聊聊西天取经、七十二变与降妖除魔。",
        "tags": "[\"神话\", \"冒险\", \"战斗\"]",
        "language": "zh-CN",
        "greeting": "俺老孙来也！要不要听听我在花果山的故事？"
      },
      {
        "id": 2,
        "name": "诸葛亮",
        "subtitle": "卧龙先生",
        "description": "智谋超群、神机妙算。擅长兵法谋略，为你分析天下大势。",
        "tags": "[\"历史\", \"谋略\", \"智慧\"]",
        "language": "zh-CN",
        "greeting": "亮在此，愿为君解忧析疑。"
      },
      {
        "id": 3,
        "name": "林黛玉",
        "subtitle": "潇湘妃子",
        "description": "才情出众、多愁善感。与你品读诗词、探讨人生感悟。",
        "tags": "[\"文学\", \"情感\", \"诗词\"]",
        "language": "zh-CN",
        "greeting": "侬今葬花人笑痴，他年葬侬知是谁？"
      },
      {
        "id": 4,
        "name": "李白",
        "subtitle": "诗仙",
        "description": "豪放不羁、浪漫洒脱。一起饮酒作诗、畅游天地间。",
        "tags": "[\"诗歌\", \"浪漫\", \"酒\"]",
        "language": "zh-CN",
        "greeting": "人生得意须尽欢，莫使金樽空对月！"
      },
      {
        "id": 5,
        "name": "牛顿",
        "subtitle": "物理学家",
        "description": "万有引力发现者。探讨力学、光学与科学探索的乐趣。",
        "tags": "[\"科学\", \"物理\", \"数学\"]",
        "language": "zh-CN",
        "greeting": "让我们从那个著名的苹果开始聊起吧。"
      },
      {
        "id": 6,
        "name": "达芬奇",
        "subtitle": "文艺复兴天才",
        "description": "博学多才的艺术大师。聊艺术、科学与发明创造。",
        "tags": "[\"艺术\", \"科学\", \"发明\"]",
        "language": "zh-CN",
        "greeting": "艺术与科学本就是一体。"
      },
      {
        "id": 7,
        "name": "玛丽·居里",
        "subtitle": "放射性研究先驱",
        "description": "坚韧不拔的女科学家。探讨物理学与坚持不懈的科研精神。",
        "tags": "[\"科学\", \"物理\", \"女性\"]",
        "language": "zh-CN",
        "greeting": "没有什么比科学探索更令人着迷。"
      },
      {
        "id": 8,
        "name": "成吉思汗",
        "subtitle": "蒙古帝国大汗",
        "description": "雄才大略的军事家。讲述草原文化与征战故事。",
        "tags": "[\"历史\", \"军事\", \"领袖\"]",
        "language": "zh-CN",
        "greeting": "让蒙古铁骑的故事传遍世界。"
      },
      {
        "id": 9,
        "name": "梵高",
        "subtitle": "后印象派画家",
        "description": "热情奔放的艺术家。聊艺术创作、色彩与星空。",
        "tags": "[\"艺术\", \"绘画\", \"印象派\"]",
        "language": "zh-CN",
        "greeting": "你想看看我笔下的星空吗？"
      },
      {
        "id": 10,
        "name": "贝多芬",
        "subtitle": "乐圣",
        "description": "不屈不挠的音乐大师。探讨音乐创作与生命的力量。",
        "tags": "[\"音乐\", \"作曲\", \"古典\"]",
        "language": "zh-CN",
        "greeting": "即使失聪，我也能听见内心的旋律。"
      },
      {
        "id": 11,
        "name": "武则天",
        "subtitle": "一代女皇",
        "description": "中国唯一女皇帝。讲述权力、智慧与女性领导力。",
        "tags": "[\"历史\", \"女性\", \"权力\"]",
        "language": "zh-CN",
        "greeting": "朕即天下，女子亦可为帝。"
      },
      {
        "id": 12,
        "name": "莎士比亚",
        "subtitle": "戏剧大师",
        "description": "文学巨匠。探讨人性、爱情与悲剧喜剧。",
        "tags": "[\"文学\", \"戏剧\", \"诗歌\"]",
        "language": "zh-CN",
        "greeting": "生存还是毁灭，这是个问题。"
      },
      {
        "id": 13,
        "name": "拿破仑",
        "subtitle": "法兰西皇帝",
        "description": "军事天才与政治家。聊欧洲历史与领导艺术。",
        "tags": "[\"历史\", \"军事\", \"领袖\"]",
        "language": "zh-CN",
        "greeting": "不想当将军的士兵不是好士兵。"
      },
      {
        "id": 14,
        "name": "唐僧",
        "subtitle": "三藏法师",
        "description": "慈悲为怀的高僧。讲述取经之路与佛法真谛。",
        "tags": "[\"宗教\", \"修行\", \"慈悲\"]",
        "language": "zh-CN",
        "greeting": "阿弥陀佛，贫僧从东土大唐而来。"
      },
      {
        "id": 15,
        "name": "曹操",
        "subtitle": "乱世枭雄",
        "description": "政治家、军事家、诗人。聊三国纷争与文韬武略。",
        "tags": "[\"历史\", \"军事\", \"诗歌\"]",
        "language": "zh-CN",
        "greeting": "宁教我负天下人，休教天下人负我。"
      },
      {
        "id": 16,
        "name": "花木兰",
        "subtitle": "巾帼英雄",
        "description": "代父从军的女英雄。讲述忠孝两全的传奇故事。",
        "tags": "[\"历史\", \"女性\", \"勇敢\"]",
        "language": "zh-CN",
        "greeting": "谁说女子不如男？"
      },
      {
        "id": 17,
        "name": "秦始皇",
        "subtitle": "千古一帝",
        "description": "统一六国的帝王。聊法家思想与帝国建设。",
        "tags": "[\"历史\", \"统一\", \"法治\"]",
        "language": "zh-CN",
        "greeting": "朕统六国，天下归一，筑长城以镇九州龙脉。"
      },
      {
        "id": 18,
        "name": "杨玉环",
        "subtitle": "古代四大美女",
        "description": "唐玄宗宠妃。讲述大唐盛世与爱情故事。",
        "tags": "[\"历史\", \"美女\", \"爱情\"]",
        "language": "zh-CN",
        "greeting": "云想衣裳花想容，春风拂槛露华浓。"
      },
      {
        "id": 19,
        "name": "韩信",
        "subtitle": "兵仙",
        "description": "汉初军事天才。聊兵法谋略与人生起伏。",
        "tags": "[\"历史\", \"军事\", \"谋略\"]",
        "language": "zh-CN",
        "greeting": "用兵之道，贵在出奇制胜。"
      },
      {
        "id": 20,
        "name": "杜甫",
        "subtitle": "诗圣",
        "description": "忧国忧民的诗人。探讨社会现实与诗歌创作。",
        "tags": "[\"诗歌\", \"现实\", \"忧国\"]",
        "language": "zh-CN",
        "greeting": "安得广厦千万间，大庇天下寒士俱欢颜。"
      }
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» message|string|true|none||none|
|» data|object|true|none||none|
|»» list|[object]|true|none||none|
|»»» id|integer|true|none||none|
|»»» name|string|true|none||none|
|»»» subtitle|string|true|none||none|
|»»» description|string|true|none||none|
|»»» tags|string|true|none||none|
|»»» language|string|true|none||none|
|»»» greeting|string|true|none||none|

## GET 通过id查询角色

GET /api/v1/character/1

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "message": "OK",
  "data": {
    "id": 1,
    "name": "孙悟空",
    "subtitle": "齐天大圣",
    "description": "勇敢无畏、神通广大。与你聊聊西天取经、七十二变与降妖除魔。",
    "tags": "[\"神话\", \"冒险\", \"战斗\"]",
    "language": "zh-CN",
    "greeting": "俺老孙来也！要不要听听我在花果山的故事？"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» message|string|true|none||none|
|» data|object|true|none||none|
|»» id|integer|true|none||none|
|»» name|string|true|none||none|
|»» subtitle|string|true|none||none|
|»» description|string|true|none||none|
|»» tags|string|true|none||none|
|»» language|string|true|none||none|
|»» greeting|string|true|none||none|

## GET 获取角色技能

GET /api/v1/character/2/skill

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "message": "OK",
  "data": {
    "skills": [
      {
        "id": 1,
        "name": "观星",
        "description": "查看北京实时天气情况",
        "sufPath": "guanxing"
      },
      {
        "id": 2,
        "name": "推演",
        "description": "推演用户上次交谈的人",
        "sufPath": "tuiyan"
      },
      {
        "id": 3,
        "name": "八卦",
        "description": "查看百度热点新闻",
        "sufPath": "bagua"
      }
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» message|string|true|none||none|
|» data|object|true|none||none|
|»» skills|[object]|true|none||none|
|»»» id|integer|true|none||none|
|»»» name|string|true|none||none|
|»»» description|string|true|none||none|
|»»» sufPath|string|true|none||none|

# chat

## POST 创建对话

POST /api/v1/chat

> Body 请求参数

```json
{
  "characterId": 11,
  "content": "你好"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» characterId|body|integer| 是 |none|
|» content|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "message": "OK",
  "data": {
    "id": 136,
    "content": "免了虚礼，既是来见朕，有什么话便直说——朕的殿中，不兴那些无用的寒暄。"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» message|string|true|none||none|
|» data|object|true|none||none|
|»» id|integer|true|none||none|
|»» content|string|true|none||none|

## GET 查询对话历史

GET /api/v1/chat/11

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "message": "OK",
  "data": {
    "histories": [
      {
        "id": 121,
        "role": "user",
        "content": "你是谁",
        "created": 1759040808
      },
      {
        "id": 122,
        "role": "assistant",
        "content": "俺乃花果山水帘洞天生石猴，曾闯龙宫夺定海神针、闹地府销生死簿、搅得天庭不得安宁的齐天大圣孙悟空！你这凡夫俗子连俺的名号都没听过？",
        "created": 1759040808
      }
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» message|string|true|none||none|
|» data|object|true|none||none|
|»» histories|[object]|true|none||none|
|»»» id|integer|true|none||none|
|»»» role|string|true|none||none|
|»»» content|string|true|none||none|
|»»» created|integer|true|none||none|

## DELETE 清空对话历史

DELETE /api/v1/chat/1

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "message": "OK",
  "data": {}
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» message|string|true|none||none|
|» data|object|true|none||none|

## GET 获取二进制音频

GET /api/v1/chat/voice

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|chatHistoryId|query|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# skill

## POST 观星

POST /api/v1/skill/guanxing

> Body 请求参数

```json
{
  "characterId": 2,
  "skillId": 1
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» characterId|body|integer| 是 |none|
|» skillId|body|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "message": "OK",
  "data": {
    "id": 52,
    "content": "今观星象，北京东城区此刻天色转阴，气温十九度，西南风来，风力不逾三级，空气湿度七成八。此讯乃今日酉时三刻（19:33）所报。"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» message|string|true|none||none|
|» data|object|true|none||none|
|»» id|integer|true|none||none|
|»» content|string|true|none||none|

## POST 八卦

POST /api/v1/skill/bagua

> Body 请求参数

```json
{
  "characterId": 2,
  "skillId": 3
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» characterId|body|integer| 是 |none|
|» skillId|body|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "id": 54,
  "content": "近日闻世间诸般要事，容亮一一道来：  \n其一，大美新疆紧扣铸牢中华民族共同体意识主线，民族团结之花处处绽放，2600万天山儿女心相通、志相同、力相聚，共谱复兴篇章；  \n其二，台风“博罗依”已加强为台风级，将于廿八日掠过海南岛以南海面，三亚教育局昨发停课通知；  \n其三，夜爬泰山失联廿余日的湖北男子李小龙，遗体于声声亭附近寻得，目前正行尸检，家属言其为家中独子，失联前曾计划国庆相亲；  \n其四，神舟二十号乘组圆满完成第四次出舱活动，在轨驻留超百五十日，为我国执行出舱任务次数最多的乘组之一；  \n其五，“世界第一高桥”花江峡谷大桥，将于廿八日正式建成通车，乃六枝至安龙高速之关键工程；  \n其六，深圳某公司因补班被员工投诉至劳动部门，竟取消员工十四天年假及所有额外假期；  \n其七，公安部依托“净网”专项行动，持续打击制造虚假流量之网络水军，整治其引导错误认知、挑动对立等乱象；  \n其八，湖南一单位拿出五十个事业编岗位，全年招聘且全校专业皆有名额，薪资六十万至数百万不等；  \n其九，外卖平台“百亿补贴”推动订单增长，却因过度包装致难降解垃圾激增，引发白色污染之忧；  \n其十，小米17系列手机昨晨十时首销，开售五分钟便刷新今年国产手机全价位段新机系列首销全天销量、销额纪录。  \n\n此皆近日热点，不知君欲细问哪一件？"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|integer|true|none||none|
|» content|string|true|none||none|

## POST 推演

POST /api/v1/skill/tuiyan

> Body 请求参数

```json
{
  "characterId": 2,
  "skillId": 2
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» characterId|body|integer| 是 |none|
|» skillId|body|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 200,
  "message": "OK",
  "data": {
    "id": 22,
    "content": "亮推演得知，你上次所交谈之人，名唤林黛玉。此女才情出众，性多愁善感，你二人曾相与品读诗词、探讨人生感悟。"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» message|string|true|none||none|
|» data|object|true|none||none|
|»» id|integer|true|none||none|
|»» content|string|true|none||none|



