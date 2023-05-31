package structural

import "fmt"

// game texture example
// func GetTexture return only unique textures. If the structure already exists, then returns it
// no duplicates are created in memory

type Texture struct {
	image string
}

func NewTexture(img string) *Texture {
	return &Texture{
		image: img,
	}
}

func (t *Texture) Render(x, y int) {
	fmt.Printf("rendering texture at (%d, %d) with image data: %s\n", x, y, t.image)
}

type TextureFactory struct {
	textures map[string]Texture
}

func (f *TextureFactory) GetTexture(image string) *Texture {
	if texture, ok := f.textures[image]; ok {
		fmt.Printf("returning existing texture %s\n", texture.image)
		return &texture
	}

	fmt.Printf("creating new texture %s\n", image)
	texture := NewTexture(image)
	f.textures[image] = *texture
	return texture
}

func RunFlyweight() {
	factory := &TextureFactory{
		textures: make(map[string]Texture),
	}

	object1 := factory.GetTexture("texture1")
	object2 := factory.GetTexture("texture2")
	object3 := factory.GetTexture("texture1")

	object1.Render(10, 20)
	object2.Render(30, 40)
	object3.Render(50, 60)
}
