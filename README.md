awsregion
====
awsのリージョン名（オレゴンなど）と略称（us-west-2など）を一覧表示するだけの機能

## Description
awsのリージョン略称が覚えられずに毎回[aws region]と検索して、AWSのDocumentから略称名を引っ張ってきた。
毎回面倒なので1コマンドで表示するために作成。

## Requirement
無し

## Usage
~~~~
$ awsregion
us-east-1      	米国東部（バージニア北部）
us-west-2      	米国西部（オレゴン）
us-west-1      	米国西部（北カリフォルニア）
eu-west-1      	欧州（アイルランド）
eu-central-1   	欧州（フランクフルト）
ap-southeast-1 	アジアパシフィック（シンガポール）
ap-northeast-1 	アジアパシフィック（東京）
ap-southeast-2 	アジアパシフィック（シドニー）
ap-northeast-2 	アジアパシフィック (ソウル)
ap-south-1     	アジアパシフィック (ムンバイ)
sa-east-1      	南米 (サンパウロ)
~~~~

## Install for Mac
~~~~
L="/usr/local/bin/awsregion" && curl -o ${L} -L https://github.com/shogomuranushi/awsregion/releases/download/v0.1/awsregion_darwin_amd64 && chmod +x ${L}
~~~~

## Install for Linux
~~~~
L="/usr/local/bin/awsregion" && curl -o ${L} -L https://github.com/shogomuranushi/awsregion/releases/download/v0.1/awsregion_linux_amd64 && chmod +x ${L}
~~~~

## Licence

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)
