package main

import "fmt"

func main() {
    str := `us-east-1	米国東部（バージニア北部）
us-west-2	米国西部（オレゴン）
us-west-1	米国西部（北カリフォルニア）
eu-west-1	欧州（アイルランド）
eu-central-1	欧州（フランクフルト）
ap-southeast-1	アジアパシフィック（シンガポール）
ap-northeast-1	アジアパシフィック（東京）
ap-southeast-2	アジアパシフィック（シドニー）
ap-northeast-2	アジアパシフィック (ソウル)
ap-south-1	アジアパシフィック (ムンバイ)
sa-east-1	南米 (サンパウロ)`

    fmt.Printf(string(str))
}
