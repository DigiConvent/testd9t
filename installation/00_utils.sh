is_installed() {
    local package_name="$1"
    if dpkg-query -W -f='${Status}' "$package_name" 2>/dev/null | grep -q "installed"; then
        return 0
    else
        return 1
    fi
}
export -f is_installed