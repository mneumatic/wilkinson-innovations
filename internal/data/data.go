package data

import "github.com/mneumantic/wilkinson-innovations/internal/models"

// TestimonalData returns array of type struct
func TestimonialData() []models.Testimonial {
	return []models.Testimonial{
		{
			Name:    "Random Person",
			ImgFile: "placeholder-avatar.png",
			ImgAlt:  "Avatar",
			Quote:   "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ad asperiores delectus ea fugit inventore libero magni obcaecati odio reiciendis unde? At, cupiditate velit! Eveniet ex minima nesciunt, recusandae repellat veniam.",
			Hidden:  false,
		},
		{
			Name:    "Random Person",
			ImgFile: "placeholder-avatar-2.png",
			ImgAlt:  "Avatar",
			Quote:   "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ad asperiores delectus ea fugit inventore libero magni obcaecati odio reiciendis unde? At, cupiditate velit! Eveniet ex minima nesciunt, recusandae repellat veniam.",
			Hidden:  true,
		},
		{
			Name:    "Random Person",
			ImgFile: "placeholder-avatar-3.png",
			ImgAlt:  "Avatar",
			Quote:   "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ad asperiores delectus ea fugit inventore libero magni obcaecati odio reiciendis unde? At, cupiditate velit! Eveniet ex minima nesciunt, recusandae repellat veniam.",
			Hidden:  true,
		},
	}
}

func ProductData() []models.Product {
	return []models.Product{
		{
			Id:      "1",
			ImgFile: "large-chart-img-1@600.webp",
			Caption: "Front Left Side",
		},
		{
			Id:      "2",
			ImgFile: "large-chart-img-2@600.webp",
			Caption: "Left Side",
		},
		{
			Id:      "3",
			ImgFile: "product-img-3.png",
			Caption: "Back Left Side",
		},
		{
			Id:      "4",
			ImgFile: "product-img-4.png",
			Caption: "Back Left Side",
		},
	}
}
