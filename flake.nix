{
  description = "Unofficial TL;DR pages client in Go";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "tldr";
          version = "0.1.0";

          src = ./.;

          vendorHash = "sha256-+cn9NiNfsdlEteXV9MPZUf3BSLF6lFwLaKwMQls3S2A=";

          ldflags = [ "-w" "-s" ];

          subPackages = [ "cmd/tldr" ];

          meta = with pkgs.lib; {
            description = "Unofficial TL;DR pages client in Go";
            homepage = "https://github.com/pauloo27/tldr";
            license = licenses.mit;
            maintainers = [ ];
            mainProgram = "tldr";
          };
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
            gnumake
          ];
        };

        apps.default = {
          type = "app";
          program = "${self.packages.${system}.default}/bin/tldr";
        };
      }
    ) // {
      homeManagerModules.default = { config, lib, pkgs, ... }:
        let
          cfg = config.programs.tldr-go;
        in {
          options.programs.tldr-go = {
            enable = lib.mkEnableOption "tldr-go client";

            viewer = lib.mkOption {
              type = lib.types.str;
              default = "less";
              description = "Viewer program for pages";
            };

            language = lib.mkOption {
              type = lib.types.str;
              default = "en";
              description = "Default language for pages";
            };
          };

          config = lib.mkIf cfg.enable {
            home.packages = [ self.packages.${pkgs.system}.default ];

            xdg.configFile."tldr/config.toml".text = ''
              viewer = "${cfg.viewer}"
              language = "${cfg.language}"
            '';
          };
        };
    };
}
