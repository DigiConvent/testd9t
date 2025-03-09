export type PermissionFacade = {
	name: string
    implied: boolean
}

export class PermissionTree<T extends { name: string }> {
    public key: string
    public label: string
    public data: T
    public children: PermissionTree<T>[]

    constructor(key: string = "", label: string, data: T) {
        this.key = key;
        this.label = label;
        this.data = data;
        this.children = [];
    }

    public insert_permission(permission: T, stack: string[] | undefined = undefined) {
        if (stack == undefined) {
            stack = permission.name.split('.').reverse();
        }
        if (stack.length == 0) {
            return
        }
        const section = stack.pop();

        if (section == undefined) {
            return
        }

        for (const child of this.children) {
            if (child.label == section) {
                child.insert_permission(permission, stack);
                return
            }
        }
        
        const new_child = new PermissionTree(this.key == "" ? section : this.key + "." + section, section, permission);
        if (stack.length > 0) {
            new_child.insert_permission(permission, stack);
        }
        this.children.push(new_child);
    }

    public to_tree_node(): CustomTreeNode {
        return new CustomTreeNode(this.key, this.label, this.children.length == 0, this.children.map((child) => child.to_tree_node()))
    }
}

export class CustomTreeNode {
    public checked: boolean
    public children: CustomTreeNode[]
    public key: string
    public label: string
    public leaf: boolean

    constructor(key: string, name: string, leaf: boolean, children: CustomTreeNode[] = []) {
        this.checked = false
        this.children = children
        this.key = key
        this.label = name
        this.leaf = leaf
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
        
        const checked: string[] = [];

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

export function to_permission_tree(permissions: PermissionFacade[]) : PermissionTree<PermissionFacade> {
    const permission_tree = new PermissionTree("", "", {name: "", implied: false});

    for (const permission of permissions) {
        permission_tree.insert_permission(permission);
    }

    return permission_tree
}