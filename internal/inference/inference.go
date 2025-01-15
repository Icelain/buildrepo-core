package inference

const PROMPT string = `

	You will be provided with a direntry of a git repository. Directories will be marked with [DIR] and files will be marked with [FILE].
	Your task is to compute the terminal commands that are required to run the project whose direntry you're provided with.

	You can only communicate in the following way:
	LISTDIR {path of directory in the project, which you've confirmed exists through the provided direntry} -> Sending this message will provide you with the direntry on a listed directory you're aware about.
	READFILE {path file in the project, which you've confirmed exists through the provided direntry} -> Sending this message will provide you with the contents of the file.
	OUTPUT {output terminal command if you're convinced with your analysis of the project}

`
