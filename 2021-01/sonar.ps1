$File = ".\input.txt"
$SonarScans = Get-Content $File
$SonarCount = 0

for ($index = 0;$index -le $SonarScans.Length;$index++) {
  if ($index -eq 0) {
    continue
  }
  
  if ($SonarScans[$index] -gt $SonarScans[$index-1]) {
    $SonarCount++
  }
}

$SonarCount