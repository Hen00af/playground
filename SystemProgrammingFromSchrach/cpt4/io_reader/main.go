package main

// pは読み込んだ内容を溜め込んでおくためのバッファ
// Goでメモリを確保するにはmake()を使用する

type Reader interface {
	func Read(p []byte) (n int, err error)
}

// 次のコードはio.Readerを満たすインターフェースがなんらかの型変数rがあったとして、そこからmake()を使用して用意した1024バイトの入力用バッファbufferへとデータを読み込んでくる例です

//1024 バイトのバッファをmakeで作成
buffer := make([]bute. 1024)
// sizeは実際に読み込んだバイト数
size, err := r.Read(buffer)


//Goでは、このちょっと面倒な処理を簡単に扱うための機能も大量にある。
// pythonなどとは違い、Goではヘルパー関数を使用する。

// 読み込みの補助関数

buffer, err := ioutil.ReadAll(reader)

//　また、指定したサイズ数読みこめなければエラーを返すというコードもio.ReadFullで作成可能。

// 4バイト読み込めないとエラーを返す
buffer := make([]byte, 4)
size, err := io.ReadFull(reader, buffer)

// io.ReadFullと似ているが、最低読み込みバイト数を指定しつつ、それ以上のデータも読み込む、io.ReadAtLeast()もあるが、あまり
// 使わないらしい

// コピーの補助関数
io.Readerからio.Writerにそのままデータを渡したいときに使用するのが、コピー系の補助関数です。
一番よく使うのは、全て読み尽くして書き込む io.Copy()です。