/**
 * 
 */
package com.lenovo.client;

import java.net.*;
import java.io.*;
import net.sf.json.JSONObject;

/**
 * @author steven
 *
 */
public class Worker {

	private static Socket socket;
    private static BufferedReader in;
    private static PrintWriter out;
    
	/**
	 * 
	 */
	public Worker() {
		// TODO Auto-generated constructor stub
	}

	/**
	 * @param args
	 */
    
    public static void main(String[] args) {
        // TODO Auto-generated method stub
        JSONObject jsonObj=new JSONObject();
        jsonObj.put("Name","liangyongs");
        jsonObj.put("Id", 31);
        String[] likes={"java","golang","clang"};
        jsonObj.put("Lks", likes);
        System.out.println("Object before sending to golang side:");
        System.out.println(jsonObj);
        try{
            socket=new Socket("127.0.0.1",50000);
            in=new BufferedReader(new InputStreamReader(socket.getInputStream()));
            out=new PrintWriter(socket.getOutputStream(),true);
            out.println(jsonObj);
            String line=in.readLine();
            System.out.println("Object read from golang side:");
            jsonObj=JSONObject.fromObject(line);
            System.out.println(jsonObj.get("Id")); 
            socket.close();
        }
        catch(IOException e){
            System.out.println(e);
        }
    }
	
	

}
