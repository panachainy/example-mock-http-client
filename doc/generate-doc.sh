find . -name "*.wsd" -print0 | while read -d $'\0' file
do
    cat $file | docker run --rm -i aplr/plantuml -tpng >${file%.*}.png
    echo "export \`$file\` done."
done
