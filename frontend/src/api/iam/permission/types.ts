import type { PermissionGroupFacade } from "../permission_group/types"
import type { UserFacade } from "../user/types"

export type PermissionFacade = {
   name: string
   implied: boolean
}

export type PermissionRead = {
   name: string
   description: string
   meta: string
   generated: boolean
   archived: boolean
}

export class PermissionTree<T extends { name: string }> {
   public key: string
   public label: string
   public data: T
   public children: PermissionTree<T>[]

   constructor(key: string = "", label: string, data: T) {
      this.key = key
      this.label = label
      this.data = data
      this.children = []
   }

   public insert_permission(permission: T, stack: string[] | undefined = undefined) {
      if (stack == undefined) {
         stack = permission.name.split(".").reverse()
      }
      if (stack.length == 0) {
         return
      }
      const section = stack.pop()

      if (section == undefined) {
         return
      }

      for (const child of this.children) {
         if (child.label == section) {
            child.insert_permission(permission, stack)
            return
         }
      }

      const new_child = new PermissionTree(
         this.key == "" ? section : this.key + "." + section,
         section,
         permission,
      )
      if (stack.length > 0) {
         new_child.insert_permission(permission, stack)
      }
      this.children.push(new_child)
   }

   public to_tree_node(parent: CustomTreeNode | null): CustomTreeNode {
      const c = new CustomTreeNode(parent, this.key, this.label, this.children.length == 0)
      const children = this.children.map((child) => child.to_tree_node(c))
      for (const child of children) {
         c.add_child(child)
      }
      return c
   }
}

export class CustomTreeNode {
   public checked: boolean
   public children: CustomTreeNode[] = []
   public key: string
   public label: string
   public leaf: boolean
   public parent: CustomTreeNode | null

   constructor(parent: CustomTreeNode | null, key: string, name: string, leaf: boolean) {
      this.parent = parent
      this.checked = false
      this.key = key
      this.label = name
      this.leaf = leaf
   }

   public uncheck(keys_to_uncheck: string[]) {
      if (keys_to_uncheck.includes(this.key)) {
         this.checked = false
         for (const child of this.children) {
            child.uncheck([child.key])
         }
      } else {
         for (const child of this.children) {
            child.uncheck(keys_to_uncheck)
         }
      }
   }
   public set_checked(keys_to_check: string[]) {
      if (keys_to_check.includes(this.key)) {
         this.checked = true
         for (const child of this.children) {
            child.set_checked([child.key])
         }
      } else {
         for (const child of this.children) {
            child.set_checked(keys_to_check)
         }
      }
   }

   public add_child(child: CustomTreeNode) {
      this.children.push(child)
      this.children.sort((child1, child2) => child1.label.localeCompare(child2.label))
   }

   public partially_checked(): boolean {
      if (this.leaf) {
         return false
      }
      const checked_children = this.children.filter((child) => child.partially_checked())
      if (checked_children.length > 0) {
         return true
      }
      return this.children.some((child) => child.fully_checked())
   }

   public get_checked(): string[] {
      if (this.leaf) {
         if (this.checked) {
            return [this.key]
         }
         return []
      }

      const checked: string[] = []

      if (this.fully_checked()) {
         return [this.key]
      }

      for (const child of this.children) {
         checked.push(...child.get_checked())
      }
      return checked
   }

   public fully_checked(): boolean {
      if (this.leaf) {
         return this.checked
      }
      return this.children.every((child) => child.fully_checked())
   }
}

export function to_permission_tree(
   permissions: PermissionFacade[],
): PermissionTree<PermissionFacade> {
   const permission_tree = new PermissionTree("", "", { name: "", implied: false })

   for (const permission of permissions) {
      permission_tree.insert_permission(permission)
   }

   return permission_tree
}

export type PermissionProfile = {
   descendants: PermissionFacade[]
   permission_groups: PermissionGroupFacade[]
   permission: PermissionRead
   users: UserFacade[]
}
