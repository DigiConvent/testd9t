export const colours = [
   "#ef4444",
   "#f59e0b",
   "#84cc16",
   "#10b981",
   "#06b6d4",
   "#3b82f6",
   "#8b5cf6",
   "#d946ef",
   "#f43f5e",
   "#f97316",
   "#eab308",
   "#22c55e",
   "#14b8a6",
   "#0ea5e9",
   "#6366f1",
   "#a855f7",
   "#ec4899",
]

export function get_colour(index: number) {
   return colours[index % colours.length]
}
