package app;

import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;

import javax.imageio.ImageIO;

public class Convert {

	private File[] listImages;

	// Construtor
	public Convert(File[] listImages) {
		this.listImages = listImages;
	}

	public void convert() throws IOException {
		int i = 1;

		for (File src : this.listImages) {

			File output = new File("//home//vagrant//workerOutput//convert_" + src.getName() + ".jpg");

			// Read 24-bit RGB input JPEG.
			BufferedImage rgbImage = ImageIO.read(src);
			int w = rgbImage.getWidth();
			int h = rgbImage.getHeight();

			// Create 8-bit gray output image from input.
			BufferedImage grayImage = new BufferedImage(w, h, BufferedImage.TYPE_BYTE_BINARY);
			int[] rgbArray = rgbImage.getRGB(0, 0, w, h, null, 0, w);
			grayImage.setRGB(0, 0, w, h, rgbArray, 0, w);

			// Save output.
			ImageIO.write(grayImage, "jpg", output);
			System.out.println("Imagem " + i + " convertida.");
			i++;

		}

	}

}
