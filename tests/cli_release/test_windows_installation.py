import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest


class WindowsInstallation(unittest.TestCase):
    def test_current_binary_is_used_for_migration_and_registration(self):
        shell = shutil.which('powershell') or shutil.which('pwsh')
        if os.name != 'nt' or not shell:
            self.skipTest('Windows PowerShell required')
        source = (Path(__file__).resolve().parents[2] / 'static/install.ps1').read_text(encoding='utf-8')
        prefix = source.split('# ── Main ─')[0]
        self.assertNotEqual(prefix, source)
        with tempfile.TemporaryDirectory(prefix='existing CLI ') as temp:
            root = Path(temp)
            binary = root / 'eigenflux.ps1'
            log = root / 'calls.txt'
            binary.write_text("if ($args[0] -eq 'version') { '99.0.0'; exit 0 }\n"
                              "Add-Content -LiteralPath $env:TEST_CALLS -Value ($args -join '|')\n"
                              "exit 0\n")
            fixture = root / 'fixture.ps1'
            fixture.write_text(prefix + """
function Invoke-RestMethod { '99.0.0' }
function Get-InvokingHost { 'codex' }
Install-Cli
if ($script:installedCliPath -ne $env:TEST_BINARY) { throw 'Wrong existing CLI path' }
Migrate-Config
& $script:installedCliPath --homedir $env:EIGENFLUX_HOME installation record --host codex --skills-target ''
if ($LASTEXITCODE -ne 0) { throw 'Registration failed' }
""", encoding='utf-8-sig')
            env = dict(os.environ, PATH=str(root) + os.pathsep + os.environ['PATH'],
                       TEST_CALLS=str(log), TEST_BINARY=str(binary),
                       EIGENFLUX_HOME=str(root / 'agent home'), USERPROFILE=str(root))
            result = subprocess.run([shell, '-NoProfile', '-ExecutionPolicy', 'Bypass',
                                     '-File', str(fixture)], env=env, capture_output=True,
                                    text=True, timeout=30)
            self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
            calls = log.read_text()
            self.assertIn('|migrate', calls)
            self.assertIn('|installation|record|--host|codex', calls)
