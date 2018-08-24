package tsFileWriter

import (
	"github.com/go_sample/src/tsfile/common/tsFileConf"
	"os"
)

/**
 * @Package Name: write
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-23 下午3:57
 * @Description:
 */

 func WriteMagic(file *os.File){
	 file.Write([]byte(tsFileConf.MAGIC_STRING))
 }