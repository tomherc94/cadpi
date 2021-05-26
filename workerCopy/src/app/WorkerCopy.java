package app;

import java.io.File;
import java.io.IOException;

public class WorkerCopy {

	public static void main(String[] args) {

		Runtime run = Runtime.getRuntime();

		long e_time = System.currentTimeMillis();

		try {
			// leio todas as imagens convertidas
			File[] filesOut = new File("//home//vagrant//workerOutput").listFiles();

			for (File src : filesOut) {

				// envia imagens processadas pela rede
				String cmd = "sshpass -p 123 scp -o StrictHostKeyChecking=no " + src.getAbsolutePath()
						+ " root@172.42.42.100:/home/vagrant/masterOutput";

				run.exec(cmd);
				try {
					Thread.sleep(50);
				} catch (Exception e) {
					e.printStackTrace();
				}

			}
			System.out.print("\n##Processo Finalizado.");

		} catch (IOException e) {
			e.printStackTrace();
			System.exit(0);
		}

		e_time = System.currentTimeMillis() - e_time;

		System.out.println("Tempo de cópia para Master: " + e_time / 1000 + "s");

	}

}
