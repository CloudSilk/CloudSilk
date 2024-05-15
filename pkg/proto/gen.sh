protoc -I . --go_out=. --go-triple_out=. *.proto

# 去掉omitempty
ls *.pb.go | xargs -n1 -IX bash -c "sed -e 's/,omitempty//' X > X.tmp && mv X{.tmp,}"

array=$(ls *.pb.go)

for a in $array; do
    protoc-go-inject-tag -input=${a}
done
