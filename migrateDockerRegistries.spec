%define debug_package   %{nil}
%define _build_id_links none
%define _name migrateDockerRegistries
%define _prefix /opt
%define _version 1.10.00
%define _rel 0
%define _arch x86_64
%define _binaryname migrateDockerRegistries

Name:       migrateDockerRegistries
Version:    %{_version}
Release:    %{_rel}
Summary:    Migrate a docker registry onto another

Group:      Docker tools
License:    GPL2.0
URL:        https://git.famillegratton.net:3000/devops/migrateDockerRegistries

Source0:    %{name}-%{_version}.tar.gz
BuildArchitectures: x86_64
BuildRequires: gcc

%description
Migrate a docker registry onto another

%prep
%autosetup

%build
cd %{_sourcedir}/%{_name}-%{_version}/src
PATH=$PATH:/opt/go/bin go build -o %{_sourcedir}/%{_binaryname} .
strip %{_sourcedir}/%{_binaryname}

%clean
rm -rf $RPM_BUILD_ROOT

%pre
exit 0

%install
install -Dpm 0755 %{_sourcedir}/%{_binaryname} %{buildroot}%{_bindir}/%{_binaryname}

%post

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{_binaryname}


%changelog
* Fri Jan 19 2024 RPM Builder <builder@famillegratton.net> 1.10.00-0
- Added missing dependencies (jean-francois@famillegratton.net)
- Automatic commit of package [migrateDockerRegistries] release [1.10.00-0].
  (builder@famillegratton.net)
- Upgraded BuildDeps (jean-francois@famillegratton.net)
- Completed shell script generation (jean-francois@famillegratton.net)
- buildDeps update (jean-francois@famillegratton.net)
- Update .gitignore (jean-francois@famillegratton.net)
- bergen-> (jean-francois@famillegratton.net)
- release bump (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)

* Fri Jan 19 2024 RPM Builder <builder@famillegratton.net> 1.10.00-0
- Upgraded BuildDeps (jean-francois@famillegratton.net)
- Completed shell script generation (jean-francois@famillegratton.net)
- buildDeps update (jean-francois@famillegratton.net)
- Update .gitignore (jean-francois@famillegratton.net)
- bergen-> (jean-francois@famillegratton.net)
- release bump (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)

* Wed Jan 17 2024 RPM Builder <builder@famillegratton.net> 1.02.00-0
- Completed minor refactoring and more verbose output (jean-
  francois@famillegratton.net)
- Completed minor refactoring and more verbose output (jean-
  francois@famillegratton.net)
- Fixed Makefile (incorrect arch) (builder@famillegratton.net)
- Specfile fix (builder@famillegratton.net)

* Wed Jan 17 2024 RPM Builder <builder@famillegratton.net> 1.01.00-0
- new package built with tito

