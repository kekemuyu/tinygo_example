tinygo build -target=bluepill -o clock.hex
tinygo build -target=bluepill -o clock.elf
arm-none-eabi-size clock.elf
