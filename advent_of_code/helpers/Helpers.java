package advent_of_code.helpers;

import java.io.BufferedReader;
import java.io.FileReader;
import java.util.ArrayList;


public class Helpers {
        
    public static ArrayList<String> readFileLineByLine(String filepath) throws Exception {
        
        BufferedReader reader;
        ArrayList <String> lines = new ArrayList <String>();
        reader = new BufferedReader(new FileReader(filepath));
        
        String line = reader.readLine();
        while (line != null) {
            lines.add(line);
            line = reader.readLine();
        }
        reader.close();

        return lines;
    
    } 
                        
}
