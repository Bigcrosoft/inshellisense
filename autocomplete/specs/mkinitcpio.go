// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["mkinitcpio"] = model.Subcommand{
		Name:        []string{"mkinitcpio"},
		Description: `Create an initial ramdisk environment`,
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Show help for mkinitcpio`,
		}, {
			Name:        []string{"--version", "-V"},
			Description: `Display version information`,
		}, {
			Name:        []string{"--addhooks", "-A"},
			Description: `Add additional hooks to the image`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "hooks",
			}},
		}, {
			Name:        []string{"--config", "-c"},
			Description: `Use config file to generate the ramdisk`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "config",
			}},
		}, {
			Name:        []string{"--generatedir", "-d"},
			Description: `Set directory as the location where the initramfs is built`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "directory",
			}},
		}, {
			Name:        []string{"--hookdir", "-D"},
			Description: `Set directory as the location where hooks will be searched for when generating the image`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "directory",
			}},
		}, {
			Name:        []string{"--generate", "-g"},
			Description: `Generate a CPIO image as filename`,
			Args: []model.Arg{{
				Name: "filename",
			}},
		}, {
			Name:        []string{"--hookhelp", "-H"},
			Description: `Output help for hookname`,
			Args: []model.Arg{{
				Name: "hookname",
			}},
		}, {
			Name:        []string{"--kernel", "-k"},
			Description: `Use kernelversion, instead of the current running kernel`,
			Args: []model.Arg{{
				Name: "kernelversion",
			}},
		}, {
			Name:        []string{"--listhooks", "-L"},
			Description: `List all available hooks`,
		}, {
			Name:        []string{"--automods", "-M"},
			Description: `Display modules found via autodetection`,
		}, {
			Name:        []string{"--nocolor", "-n"},
			Description: `Disable color output`,
		}, {
			Name:        []string{"--uefi", "-U"},
			Description: `Generate a UEFI executable as filename`,
			Args: []model.Arg{{
				Name: "filename",
			}},
		}, {
			Name:        []string{"--allpresets", "-P"},
			Description: `Process all presets contained in /etc/mkinitcpio.d`,
		}, {
			Name:        []string{"--preset", "-p"},
			Description: `Build initramfs image(s) according to specified preset`,
			Args: []model.Arg{{
				Name:      "preset",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"--moduleroot", "-r"},
			Description: `Specifies the root directory to find modules in`,
			Args: []model.Arg{{
				Name: "root",
			}},
		}, {
			Name:        []string{"--skiphooks", "-S"},
			Description: `Skip hooks when generating the image`,
			Args: []model.Arg{{
				Name: "hooks",
			}},
		}, {
			Name:        []string{"--save", "-s"},
			Description: `Saves the build directory for the initial ramdisk`,
		}, {
			Name:        []string{"--builddir", "-t"},
			Description: `Use tmpdir as the temporary build directory instead of /tmp`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFolders},
				Name:      "tmpdir",
			}},
		}, {
			Name:        []string{"--verbose", "-v"},
			Description: `Verbose output`,
		}, {
			Name:        []string{"--compress", "-z"},
			Description: `Override the compression method with the compress program`,
			Args: []model.Arg{{
				Name: "compress",
			}},
		}, {
			Name:        []string{"--cmdline"},
			Description: `Use kernel command line with UEFI executable`,
			Args: []model.Arg{{
				Name: "filename",
			}},
		}, {
			Name:        []string{"--splash"},
			Description: `UEFI executables can show a bitmap file on boot`,
			Args: []model.Arg{{
				Name: "filename",
			}},
		}, {
			Name:        []string{"--uefistub"},
			Description: `UEFI stub image used for UEFI executable generation`,
			Args: []model.Arg{{
				Name: "filename",
			}},
		}, {
			Name:        []string{"--kernelimage"},
			Description: `Include a kernel image for the UEFI executable`,
			Args: []model.Arg{{
				Name: "filename",
			}},
		}, {
			Name:        []string{"--microcode"},
			Description: `Include microcode into the UEFI executable`,
			Args: []model.Arg{{
				Name: "filename",
			}},
		}, {
			Name:        []string{"--osrelease"},
			Description: `Include a os-release file for the UEFI executable`,
			Args: []model.Arg{{
				Name: "filename",
			}},
		}},
	}
}