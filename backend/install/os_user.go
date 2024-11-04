package install

import installation_utils "github.com/DigiConvent/testd9t/install/utils"

func DeleteUser() {
	installation_utils.Execute("userdel digiconvent", false)
	installation_utils.Execute("groupdel digiconvent_group", false)
	installation_utils.Execute("rm -rf /home/digiconvent", false)
}

func CreateUser() {
	installation_utils.Execute("groupadd digiconvent_group", false)
	installation_utils.Execute("useradd -g digiconvent_group digiconvent", false)
	installation_utils.Execute("usermod -aG sudo digiconvent", false)
	installation_utils.Execute("mkdir /home/digiconvent", false)
	installation_utils.Execute("chown -R digiconvent:digiconvent_group /home/digiconvent", false)
}
