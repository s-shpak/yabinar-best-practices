#! /bin/bash
set -e

cd "${1}"

main() {
    if go mod tidy -v 2>&1 | grep -q 'updates to go.mod needed'; then
        exit 1
    fi

    local go_mod_file="go.mod"
    declare go_mod_status

    if [[ -e $go_mod_file ]]; then
        set +e
        git diff --exit-code $go_mod_file &> /dev/null
        go_mod_status=$?
        set -e
    else
        echo "файл '$go_mod_file' не найден. Инициализируйте go-модуль в корне проекта при помощи 'go mod init'"
        exit 1
    fi

    if [ $go_mod_status -ne 0 ]; then
        echo "'$go_mod_file' был обновлен, добавьте его в коммит"
        exit 1
    fi

    local go_sum_file="go.sum"
    local go_sum_status=0

    if [[ -e $go_sum_file ]]; then
        set +e
        git diff --exit-code $go_sum_file &> /dev/null
        go_sum_status=$?
        set -e
    fi

    if [ $go_sum_status -ne 0 ]; then
        echo "'$go_sum_file' был обновлен, добавьте его в коммит"
        exit 1
    fi
}

main
