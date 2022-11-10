This is a replication of a stack overflow panic that occurs when trying to run a mutation
against a Dgraph instance using Go http client.

```console
docker run --network=host -d -p 8080:8080 dgraph/standalone:v22.0.0
curl -X POST localhost:8080/admin/schema --data-binary '@schema.graphql'

go run main.go
```


OUTPUT:

```
runtime: goroutine stack exceeds 1000000000-byte limit
runtime: sp=0xc0203003c0 stack=[0xc020300000, 0xc040300000]
fatal error: stack overflow

runtime stack:
runtime.throw({0x6aac3f?, 0x870ee0?})
	/usr/local/go/src/runtime/panic.go:1047 +0x5d fp=0x7f10021dcc18 sp=0x7f10021dcbe8 pc=0x43665d
runtime.newstack()
	/usr/local/go/src/runtime/stack.go:1103 +0x5cc fp=0x7f10021dcdd0 sp=0x7f10021dcc18 pc=0x44f6ac
runtime.morestack()
	/usr/local/go/src/runtime/asm_amd64.s:570 +0x8b fp=0x7f10021dcdd8 sp=0x7f10021dcdd0 pc=0x463c0b

goroutine 1 [running]:
runtime.heapBitsSetType(0xc00b0d54f0?, 0x10?, 0x10?, 0x65de80?)
	/usr/local/go/src/runtime/mbitmap.go:844 +0xbac fp=0xc0203003d0 sp=0xc0203003c8 pc=0x4170cc
runtime.mallocgc(0x10, 0x65de80, 0x1)
	/usr/local/go/src/runtime/malloc.go:1050 +0x64d fp=0xc020300448 sp=0xc0203003d0 pc=0x40e70d
runtime.growslice(0x65de80, {0x0?, 0xc00b0d82a8?, 0x2?}, 0x0?)
	/usr/local/go/src/runtime/slice.go:290 +0x4ea fp=0xc0203004b0 sp=0xc020300448 pc=0x44d5ca
github.com/shurcooL/graphql/ident.ParseMixedCaps({0x64fe7a?, 0x2?})
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/ident/ident.go:45 +0x549 fp=0xc020300630 sp=0xc0203004b0 pc=0x640ba9
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x684120}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:120 +0x291 fp=0xc020300730 sp=0xc020300630 pc=0x642511
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020300830 sp=0xc020300730 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020300930 sp=0xc020300830 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020300a30 sp=0xc020300930 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020300b30 sp=0xc020300a30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020300c30 sp=0xc020300b30 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020300d30 sp=0xc020300c30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020300e30 sp=0xc020300d30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020300f30 sp=0xc020300e30 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020301030 sp=0xc020300f30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020301130 sp=0xc020301030 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020301230 sp=0xc020301130 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020301330 sp=0xc020301230 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020301430 sp=0xc020301330 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020301530 sp=0xc020301430 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020301630 sp=0xc020301530 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020301730 sp=0xc020301630 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020301830 sp=0xc020301730 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020301930 sp=0xc020301830 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020301a30 sp=0xc020301930 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020301b30 sp=0xc020301a30 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020301c30 sp=0xc020301b30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020301d30 sp=0xc020301c30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020301e30 sp=0xc020301d30 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020301f30 sp=0xc020301e30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020302030 sp=0xc020301f30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020302130 sp=0xc020302030 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020302230 sp=0xc020302130 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020302330 sp=0xc020302230 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020302430 sp=0xc020302330 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020302530 sp=0xc020302430 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020302630 sp=0xc020302530 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020302730 sp=0xc020302630 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020302830 sp=0xc020302730 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020302930 sp=0xc020302830 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020302a30 sp=0xc020302930 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020302b30 sp=0xc020302a30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020302c30 sp=0xc020302b30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020302d30 sp=0xc020302c30 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020302e30 sp=0xc020302d30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020302f30 sp=0xc020302e30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020303030 sp=0xc020302f30 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020303130 sp=0xc020303030 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020303230 sp=0xc020303130 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020303330 sp=0xc020303230 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020303430 sp=0xc020303330 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020303530 sp=0xc020303430 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020303630 sp=0xc020303530 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020303730 sp=0xc020303630 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020303830 sp=0xc020303730 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020303930 sp=0xc020303830 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020303a30 sp=0xc020303930 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020303b30 sp=0xc020303a30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020303c30 sp=0xc020303b30 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020303d30 sp=0xc020303c30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020303e30 sp=0xc020303d30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020303f30 sp=0xc020303e30 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020304030 sp=0xc020303f30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020304130 sp=0xc020304030 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020304230 sp=0xc020304130 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020304330 sp=0xc020304230 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020304430 sp=0xc020304330 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020304530 sp=0xc020304430 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020304630 sp=0xc020304530 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020304730 sp=0xc020304630 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020304830 sp=0xc020304730 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020304930 sp=0xc020304830 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020304a30 sp=0xc020304930 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020304b30 sp=0xc020304a30 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020304c30 sp=0xc020304b30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020304d30 sp=0xc020304c30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020304e30 sp=0xc020304d30 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020304f30 sp=0xc020304e30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020305030 sp=0xc020304f30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020305130 sp=0xc020305030 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020305230 sp=0xc020305130 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020305330 sp=0xc020305230 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020305430 sp=0xc020305330 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020305530 sp=0xc020305430 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020305630 sp=0xc020305530 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020305730 sp=0xc020305630 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020305830 sp=0xc020305730 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020305930 sp=0xc020305830 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020305a30 sp=0xc020305930 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020305b30 sp=0xc020305a30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020305c30 sp=0xc020305b30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020305d30 sp=0xc020305c30 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020305e30 sp=0xc020305d30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020305f30 sp=0xc020305e30 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020306030 sp=0xc020305f30 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020306130 sp=0xc020306030 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020306230 sp=0xc020306130 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020306330 sp=0xc020306230 pc=0x6423b0
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x6841e0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020306430 sp=0xc020306330 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x690aa0}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:123 +0x16b fp=0xc020306530 sp=0xc020306430 pc=0x6423eb
github.com/shurcooL/graphql.writeQuery({0x722d20, 0xc0000a4cf0}, {0x7276c0, 0x65a980}, 0x0)
	/home/nmoraitis/go/pkg/mod/github.com/shurcoo!l/graphql@v0.0.0-20220606043923-3cf50f8a0a29/query.go:100 +0x130 fp=0xc020306630 sp=0xc020306530 pc=0x6423b0
...additional frames elided...

goroutine 2 [force gc (idle)]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/local/go/src/runtime/proc.go:363 +0xd6 fp=0xc000052fb0 sp=0xc000052f90 pc=0x439276
runtime.goparkunlock(...)
	/usr/local/go/src/runtime/proc.go:369
runtime.forcegchelper()
	/usr/local/go/src/runtime/proc.go:302 +0xad fp=0xc000052fe0 sp=0xc000052fb0 pc=0x43910d
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1594 +0x1 fp=0xc000052fe8 sp=0xc000052fe0 pc=0x465ce1
created by runtime.init.6
	/usr/local/go/src/runtime/proc.go:290 +0x25

goroutine 3 [GC sweep wait]:
runtime.gopark(0x1?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/local/go/src/runtime/proc.go:363 +0xd6 fp=0xc000053790 sp=0xc000053770 pc=0x439276
runtime.goparkunlock(...)
	/usr/local/go/src/runtime/proc.go:369
runtime.bgsweep(0x0?)
	/usr/local/go/src/runtime/mgcsweep.go:297 +0xd7 fp=0xc0000537c8 sp=0xc000053790 pc=0x425db7
runtime.gcenable.func1()
	/usr/local/go/src/runtime/mgc.go:178 +0x26 fp=0xc0000537e0 sp=0xc0000537c8 pc=0x41ac26
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1594 +0x1 fp=0xc0000537e8 sp=0xc0000537e0 pc=0x465ce1
created by runtime.gcenable
	/usr/local/go/src/runtime/mgc.go:178 +0x6b

goroutine 4 [GC scavenge wait]:
runtime.gopark(0x102d21534c9c?, 0xad0cab?, 0x0?, 0x0?, 0x0?)
	/usr/local/go/src/runtime/proc.go:363 +0xd6 fp=0xc000053f70 sp=0xc000053f50 pc=0x439276
runtime.goparkunlock(...)
	/usr/local/go/src/runtime/proc.go:369
runtime.(*scavengerState).park(0x8a4e40)
	/usr/local/go/src/runtime/mgcscavenge.go:389 +0x53 fp=0xc000053fa0 sp=0xc000053f70 pc=0x423e13
runtime.bgscavenge(0x0?)
	/usr/local/go/src/runtime/mgcscavenge.go:622 +0x65 fp=0xc000053fc8 sp=0xc000053fa0 pc=0x424405
runtime.gcenable.func2()
	/usr/local/go/src/runtime/mgc.go:179 +0x26 fp=0xc000053fe0 sp=0xc000053fc8 pc=0x41abc6
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1594 +0x1 fp=0xc000053fe8 sp=0xc000053fe0 pc=0x465ce1
created by runtime.gcenable
	/usr/local/go/src/runtime/mgc.go:179 +0xaa

goroutine 18 [finalizer wait]:
runtime.gopark(0x8a53c0?, 0xc00008a4e0?, 0x0?, 0x0?, 0xc000052770?)
	/usr/local/go/src/runtime/proc.go:363 +0xd6 fp=0xc000052628 sp=0xc000052608 pc=0x439276
runtime.goparkunlock(...)
	/usr/local/go/src/runtime/proc.go:369
runtime.runfinq()
	/usr/local/go/src/runtime/mfinal.go:180 +0x10f fp=0xc0000527e0 sp=0xc000052628 pc=0x419d2f
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1594 +0x1 fp=0xc0000527e8 sp=0xc0000527e0 pc=0x465ce1
created by runtime.createfing
	/usr/local/go/src/runtime/mfinal.go:157 +0x45

goroutine 19 [GC worker (idle)]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/local/go/src/runtime/proc.go:363 +0xd6 fp=0xc00004e750 sp=0xc00004e730 pc=0x439276
runtime.gcBgMarkWorker()
	/usr/local/go/src/runtime/mgc.go:1235 +0xf1 fp=0xc00004e7e0 sp=0xc00004e750 pc=0x41cb71
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1594 +0x1 fp=0xc00004e7e8 sp=0xc00004e7e0 pc=0x465ce1
created by runtime.gcBgMarkStartWorkers
	/usr/local/go/src/runtime/mgc.go:1159 +0x25

goroutine 20 [GC worker (idle)]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/local/go/src/runtime/proc.go:363 +0xd6 fp=0xc00004ef50 sp=0xc00004ef30 pc=0x439276
runtime.gcBgMarkWorker()
	/usr/local/go/src/runtime/mgc.go:1235 +0xf1 fp=0xc00004efe0 sp=0xc00004ef50 pc=0x41cb71
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1594 +0x1 fp=0xc00004efe8 sp=0xc00004efe0 pc=0x465ce1
created by runtime.gcBgMarkStartWorkers
	/usr/local/go/src/runtime/mgc.go:1159 +0x25

goroutine 34 [GC worker (idle)]:
runtime.gopark(0x102dab3be457?, 0x65a980?, 0x53?, 0x2e?, 0x1a?)
	/usr/local/go/src/runtime/proc.go:363 +0xd6 fp=0xc000488750 sp=0xc000488730 pc=0x439276
runtime.gcBgMarkWorker()
	/usr/local/go/src/runtime/mgc.go:1235 +0xf1 fp=0xc0004887e0 sp=0xc000488750 pc=0x41cb71
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1594 +0x1 fp=0xc0004887e8 sp=0xc0004887e0 pc=0x465ce1
created by runtime.gcBgMarkStartWorkers
	/usr/local/go/src/runtime/mgc.go:1159 +0x25

goroutine 35 [GC worker (idle)]:
runtime.gopark(0x102dab3bee9f?, 0x0?, 0xc0?, 0x76?, 0x65a980?)
	/usr/local/go/src/runtime/proc.go:363 +0xd6 fp=0xc000488f50 sp=0xc000488f30 pc=0x439276
runtime.gcBgMarkWorker()
	/usr/local/go/src/runtime/mgc.go:1235 +0xf1 fp=0xc000488fe0 sp=0xc000488f50 pc=0x41cb71
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1594 +0x1 fp=0xc000488fe8 sp=0xc000488fe0 pc=0x465ce1
created by runtime.gcBgMarkStartWorkers
	/usr/local/go/src/runtime/mgc.go:1159 +0x25

goroutine 36 [GC worker (idle)]:
runtime.gopark(0x8d5c20?, 0x3?, 0xe3?, 0xa6?, 0x1d?)
	/usr/local/go/src/runtime/proc.go:363 +0xd6 fp=0xc000489750 sp=0xc000489730 pc=0x439276
runtime.gcBgMarkWorker()
	/usr/local/go/src/runtime/mgc.go:1235 +0xf1 fp=0xc0004897e0 sp=0xc000489750 pc=0x41cb71
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1594 +0x1 fp=0xc0004897e8 sp=0xc0004897e0 pc=0x465ce1
created by runtime.gcBgMarkStartWorkers
	/usr/local/go/src/runtime/mgc.go:1159 +0x25

goroutine 37 [GC worker (idle)]:
runtime.gopark(0x102db0ea5cca?, 0x1?, 0xf1?, 0x4c?, 0x1a?)
	/usr/local/go/src/runtime/proc.go:363 +0xd6 fp=0xc000489f50 sp=0xc000489f30 pc=0x439276
runtime.gcBgMarkWorker()
	/usr/local/go/src/runtime/mgc.go:1235 +0xf1 fp=0xc000489fe0 sp=0xc000489f50 pc=0x41cb71
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1594 +0x1 fp=0xc000489fe8 sp=0xc000489fe0 pc=0x465ce1
created by runtime.gcBgMarkStartWorkers
	/usr/local/go/src/runtime/mgc.go:1159 +0x25

goroutine 5 [GC worker (idle)]:
runtime.gopark(0x102d6167c718?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/local/go/src/runtime/proc.go:363 +0xd6 fp=0xc000054750 sp=0xc000054730 pc=0x439276
runtime.gcBgMarkWorker()
	/usr/local/go/src/runtime/mgc.go:1235 +0xf1 fp=0xc0000547e0 sp=0xc000054750 pc=0x41cb71
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1594 +0x1 fp=0xc0000547e8 sp=0xc0000547e0 pc=0x465ce1
created by runtime.gcBgMarkStartWorkers
	/usr/local/go/src/runtime/mgc.go:1159 +0x25

goroutine 6 [GC worker (idle)]:
runtime.gopark(0x102dab3be8d7?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/local/go/src/runtime/proc.go:363 +0xd6 fp=0xc000054f50 sp=0xc000054f30 pc=0x439276
runtime.gcBgMarkWorker()
	/usr/local/go/src/runtime/mgc.go:1235 +0xf1 fp=0xc000054fe0 sp=0xc000054f50 pc=0x41cb71
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1594 +0x1 fp=0xc000054fe8 sp=0xc000054fe0 pc=0x465ce1
created by runtime.gcBgMarkStartWorkers
	/usr/local/go/src/runtime/mgc.go:1159 +0x25
exit status 2
```