# Parse

通过 `Parse` 解析 JWT 字符串

运行结果：
```sh
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJlbHlzaWEiLCJpc3MiOiJweWtlIiwic3ViIjoicHlrZWx5c2lhLmdpdGh1Yi5pbyJ9.xy6KmsE904PI7Y7yCGeJEZmf6jOuSGNe5kWcmtvb_F4 

 &{eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJlbHlzaWEiLCJpc3MiOiJweWtlIiwic3ViIjoicHlrZWx5c2lhLmdpdGh1Yi5pbyJ9.xy6KmsE904PI7Y7yCGeJEZmf6jOuSGNe5kWcmtvb_F4 0xc0000080a8 map[alg:HS256 typ:JWT] map[aud:elysia iss:pyke sub:pykelysia.github.io] [199 46 138 154 193 61 211 131 200 237 142 242 8 103 137 17 153 159 234 51 174 72 99 94 230 69 156 154 219 219 252 94] true}

 map[aud:elysia iss:pyke sub:pykelysia.github.io]
```