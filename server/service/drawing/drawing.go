package drawing

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//DrawTextInfo 图片绘字信息
type DrawTextInfo struct {
	Text   string
	X      int
	Y      int
	brunch *TextBrush
}

//TextBrush 字体刷子 相关
type TextBrush struct {
	FontType  *truetype.Font
	FontSize  float64
	FontColor *image.Uniform
}

//NewTextBrush 新生成笔刷
func NewTextBrush(FontFilePath string, FontSize float64, FontColor *image.Uniform) (*TextBrush, error) {
	fontFile, err := ioutil.ReadFile(FontFilePath)
	if err != nil {
		return nil, err
	}
	fontType, err := truetype.Parse(fontFile)
	if err != nil {
		return nil, err
	}

	return &TextBrush{FontType: fontType, FontSize: FontSize, FontColor: FontColor}, nil
}

func NewFontColor(R, G, B, A uint8) *image.Uniform {
	return image.NewUniform(color.RGBA{
		R: R,
		G: G,
		B: B,
		A: A,
	})
}

//Image2RGBA Image2RGBA
func Image2RGBA(img image.Image) *image.RGBA {

	baseSrcBounds := img.Bounds().Max

	newWidth := baseSrcBounds.X
	newHeight := baseSrcBounds.Y

	des := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight)) // 底板
	//首先将一个图片信息存入jpg
	draw.Draw(des, des.Bounds(), img, img.Bounds().Min, draw.Over)

	return des
}

//DrawStringOnImageAndSave 图片上写文字
func DrawStringOnImageAndSave(imageData []byte, infos []*DrawTextInfo, outPath string) (err error) {
	//判断图片类型
	var backgroud image.Image
	filetype := http.DetectContentType(imageData)
	switch filetype {
	case "image/jpeg", "image/jpg":
		backgroud, err = jpeg.Decode(bytes.NewReader(imageData))
		if err != nil {
			fmt.Println("jpeg error")
			return err
		}

	case "image/gif":
		backgroud, err = gif.Decode(bytes.NewReader(imageData))
		if err != nil {
			return err
		}

	case "image/png":
		backgroud, err = png.Decode(bytes.NewReader(imageData))
		if err != nil {
			return err
		}
	default:
		return err
	}
	des := Image2RGBA(backgroud)

	//新建笔刷
	//textBrush, _ := NewTextBrush("ttf/arial.ttf", 50, image.Black)
	for _, info := range infos {
		//Px Py 绘图开始坐标 text要绘制的文字
		//调整颜色
		c := freetype.NewContext()
		c.SetDPI(72)
		c.SetFont(info.brunch.FontType)
		c.SetHinting(font.HintingFull)
		c.SetFontSize(info.brunch.FontSize)
		c.SetClip(des.Bounds())
		c.SetDst(des)
		c.SetSrc(info.brunch.FontColor)
		c.DrawString(info.Text, freetype.Pt(info.X, info.Y))
	}

	//保存图片
	fSave, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer fSave.Close()

	err = png.Encode(fSave, des)

	if err != nil {
		return err
	}
	return nil
}
