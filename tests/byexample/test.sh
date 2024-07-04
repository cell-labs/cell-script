source ../../install.sh

directory="."
suffix=".cell"

# Iterate over all files with the given suffix
for file in "$directory"/*"$suffix"
do
    # Check if the file exists (in case no files match the pattern)
    if [ -e "$file" ]; then
        echo "Processing $file"
        filename=$(basename -- "$file")
        # Remove the suffix to get the name without suffix
        name_without_suffix="${filename%$suffix}"
        # Your code to process the file goes here
        cell $file -t riscv
        ckb-debugger --bin $name_without_suffix
    fi
done