export type Page<T> = {
    items: T[],
    page: number,
    items_per_page: number,
    total_items: number
}