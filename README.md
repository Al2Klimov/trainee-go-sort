# trainee-sort

## About

**trainee-sort** is a from scratch developed program which sorts text files. There are
various parameters you can use to filter or individualize the output of your to sorted file.
Seen like this it's an imitation of the original sort command line command.

## Parameters

| Parameter    | Usage | 
| ------------- |-------------|
| **-n**      | Sort numbers beginning from the lowest to the highest.
| **-o**      | File in which the output is saved. If the file doesn't exist it'll created.  
| **-u**  | Allows you to not print equal values. The value is only printed once no matter how often it appears.    
| **-b**     | Ignore the leading blanks.
| **-f**     | Fold lower case characters to upper case ones.     
| **-R**   | Shuffles the output.    
| **-r**      | Prints the output in reverse order.
| **-qsort**   | Uses the quick sort algorithm.   

## Usage Examples

In the following you'll see a quick example of how to use the sort command. Here is what our text file looks like:

toSort.txt:

                3
                5
                3
                6
                8
                5
                2
                1
                4

Now we can sort it with `./sort.linux-amd64 toSort.txt `  

The output is:

                1
                2
                3
                3
                4
                5
                5
                6
                8
                
You can also sort with `./sort.linux-amd64 < toSort.txt `. It does not change the output though.  

To save the output in a file you can either use the -o parameter or just add a "> output.txt" at the end. 
Your command should then look like this `./sort.linux-amd64 < toSort.txt > output.txt ` .
