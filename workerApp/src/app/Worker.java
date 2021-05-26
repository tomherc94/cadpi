package app;

import java.io.File;
import java.io.IOException;

public class Worker {

	public static void main(String[] args) {

		// leio todas as imagens
		File[] files = new File("//home//vagrant//workerInput").listFiles();

		// instancio o objeto de conversão passando a lista de imagens como argumento
		Convert convert = new Convert(files);

		long e_time = System.currentTimeMillis();

		try {
			convert.convert();
		} catch (IOException e) {
			e.printStackTrace();
			System.exit(0);
		}

		e_time = System.currentTimeMillis() - e_time;

		System.out.println("Tempo de processamento de imagens: " + e_time / 1000 + "s");

	}

}
