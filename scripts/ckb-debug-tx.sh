cmd="ckb-debugger --tx-file dump.json \
    --cell-index $1 \
    --cell-type input \
    --script-group-type type \
    --bin a.out"
echo $cmd
$cmd

