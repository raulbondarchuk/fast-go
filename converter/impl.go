package converter

// Convert() converts the image to the desired format and saves it to the directory
// This method is a public interface for converting images using the configuration
// specified in the ImageConfig struct. It calls the private method convertImage
// which handles the actual conversion process.
func (c *ImageConfig) Convert() (string, error) {
	// Calls the internal convertImage method to perform the conversion
	return c.convertImage()
}

// Delete() deletes the image from the directory
// This method provides a public interface to delete an image file from the
// specified directory. It accepts an optional file path parameter. If no
// path is provided, it defaults to using the FileName and DirToStorage
// from the ImageConfig struct.
func (c *ImageConfig) Delete(reqFilePath ...string) error {
	// Calls the internal deleteImage method to perform the deletion
	return c.deleteImage(reqFilePath...)
}

// Convert() handles logo upload, resizing with quality strategies, and saves it in WebP format.
// This method is a public interface for processing logos using the configuration
// specified in the LogoConfig struct. It calls the private method processLogo
// which handles the actual processing and conversion to WebP format.
func (c *LogoConfig) Convert() (string, error) {
	// Calls the internal processLogo method to perform the logo processing
	return c.processLogo()
}

// ConvertVideo is a public method to convert video using the VideoConfig settings
// This method is a public interface for converting videos using the configuration
// specified in the VideoConfig struct. It calls the private method processVideo
// which handles the actual conversion process.
func (c *VideoConfig) Convert() (string, error) {
	return c.processVideo()
}

// DeleteVideo is a public method to delete video from the directory
// This method provides a public interface to delete a video file from the
// specified directory. It accepts an optional file path parameter. If no
// path is provided, it defaults to using the FileName and DirToStorage
// from the VideoConfig struct.
func (c *VideoConfig) Delete(reqFilePath ...string) error {
	return c.deleteVideo(reqFilePath...)
}
