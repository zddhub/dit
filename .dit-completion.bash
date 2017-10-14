_dit_complete()
{
    local cur_word prev_word type_list

    # COMP_WORDS is an array of words in the current command line.
    # COMP_CWORD is the index of the current word (the one the cursor is
    # in). So COMP_WORDS[COMP_CWORD] is the current word; we also record
    # the previous word here, although this specific script doesn't
    # use it yet.
    cur_word="${COMP_WORDS[COMP_CWORD]}"
    prev_word="${COMP_WORDS[COMP_CWORD-1]}"

    # Ask dit to generate a list of sub commands it supports
    sub_commands="init add commit cat-file"

    case ${prev_word} in
        dit)
            COMPREPLY=( $(compgen -W "${sub_commands}" -- ${cur_word}) )
            return 0
            ;;
        cat-file)
            COMPREPLY=( $(compgen -W "-p -s -t -h" -- ${cur_word}) )
            return 0
            ;;
    esac
    return 0
}

complete -o bashdefault -o nospace -F _dit_complete dit

