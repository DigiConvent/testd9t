import type { PermissionGroupFacade } from "../iam/permission_group/types"

export type CustomNode<T> = {
   key: string
   data: T
   children: CustomNode<T>[]
   selectable?: boolean
   styleClass?: string
}

type HasChildren = {
   id: string
   children: string[]
}

export function to_has_parent<T extends HasChildren>(data: T[]): (T | HasParent)[] {
   const mapped_data: Map<string, T | HasParent> = new Map()
   for (const entry of data) {
      mapped_data.set(entry.id, { ...entry, parent: "" })
   }
   for (const entry of data) {
      for (const child of entry.children) {
         mapped_data.set(child, {
            ...mapped_data.get(child),
            id: child,
            parent: entry.id,
         })
      }
   }
   return [...mapped_data.values()]
}

type HasParent = {
   id: string
   parent?: string
}

export function create_tree_using_parent<T extends HasParent>(leaf: T, data: T[]): CustomNode<T> {
   const node: CustomNode<T> = {
      selectable: true,
      styleClass: "",
      key: leaf.id,
      data: leaf,
      children: [],
   }
   for (const entry of data) {
      if (entry.parent == leaf.id) {
         node.children.push(create_tree_using_parent(entry, data))
      }
   }
   return node
}

export function create_tree_using_children<T extends HasChildren>(
   leaf: T,
   data: T[],
): CustomNode<T> {
   const node: CustomNode<T> = {
      key: leaf.id,
      data: leaf,
      children: [],
   }
   for (const entry of data) {
      if (entry.children.includes(leaf.id)) {
         node.children.push(create_tree_using_children(entry, data))
      }
   }
   return node
}

export const create_tree_from_permission_group_facades = (
   leaf: PermissionGroupFacade,
   data: PermissionGroupFacade[],
): CustomNode<PermissionGroupFacade> => {
   const root: CustomNode<PermissionGroupFacade> = {
      key: leaf.id,
      data: leaf,
      children: [],
   }
   for (const entry of data) {
      if (entry.parent == leaf.id) {
         root.children.push(create_tree_from_permission_group_facades(entry, data))
      }
   }
   return root
}

export function inventory(
   node: CustomNode<any>,
   key: string,
   options: string[],
): { total: number; other: number; [key: string]: number } {
   const result: { total: number; other: number; [key: string]: number } = {
      total: node.children.length,
      other: 0,
   }
   for (const option of options) {
      result[option] = 0
      for (const child of node.children) {
         if (child.data[key] == option) result[option]++
      }
   }

   for (const child of node.children) {
      const i = inventory(child, key, options)
      result.total += i.total
      for (const option of options) {
         result[option] += i[option]
      }
   }

   result.other = result.total
   for (const option of options) {
      result.other -= result[option]
   }
   return result
}
