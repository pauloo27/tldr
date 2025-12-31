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
    );
}
