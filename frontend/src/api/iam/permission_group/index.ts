import add_user from "./add_user";
import create_permission_group from "./create";
import delete_permission_group from "./delete";
import get_permission_group from "./get";
import get_permission_group_profile from "./get_profile";
import list_permission_groups from "./list";
import set_permission_group_parent from "./set_parent";
import update_permission_group from "./update";

export const permission_group = {
    add_user: add_user,
    create: create_permission_group,
    delete: delete_permission_group,
    get: get_permission_group,
    get_profile: get_permission_group_profile,
    list: list_permission_groups,
    set_parent: set_permission_group_parent,
    update: update_permission_group,
}