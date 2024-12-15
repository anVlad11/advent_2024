#!/bin/bash

# Set variables
INPUT_PATTERN="bmp/step_%d.bmp"
STEP_TO_REPEAT=7083
FRAMES_TO_REPEAT=120
FONT_FILE="/System/Library/Fonts/Supplemental/Menlo.ttc"
OUTPUT_VIDEO="final_output.mp4"

# Generate part1.mp4 (steps before STEP_TO_REPEAT)
ffmpeg -framerate 60 -i $INPUT_PATTERN -vf "
pad=ceil(iw/2)*2:ceil(ih/2)*2:0:0,
scale=iw*5:ih*5:flags=neighbor,
select='lte(n,${STEP_TO_REPEAT}-1)',
drawtext=fontfile=${FONT_FILE}:
text='Step %{eif\:n\:d}':
x=w-tw-10:y=10:
fontsize=24:
fontcolor=white:
box=1:boxcolor=black@0.5" -fps_mode vfr -c:v libx264 -pix_fmt yuv420p part1.mp4

# Generate part2.mp4 (STEP_TO_REPEAT repeated for FRAMES_TO_REPEAT frames)
ffmpeg -loop 1 -i bmp/step_${STEP_TO_REPEAT}.bmp -vf "
pad=ceil(iw/2)*2:ceil(ih/2)*2:0:0,
scale=iw*5:ih*5:flags=neighbor,
drawtext=fontfile=${FONT_FILE}:
text='Step ${STEP_TO_REPEAT}':
x=w-tw-10:y=10:
fontsize=24:
fontcolor=green:
box=1:boxcolor=black@0.5" -t $(echo "scale=2; $FRAMES_TO_REPEAT / 60" | bc) -r 60 -c:v libx264 -pix_fmt yuv420p part2_final.mp4

# Generate part3.mp4 (steps after STEP_TO_REPEAT)
ffmpeg -framerate 60 -i $INPUT_PATTERN -vf "
pad=ceil(iw/2)*2:ceil(ih/2)*2:0:0,
scale=iw*5:ih*5:flags=neighbor,
select='gte(n,${STEP_TO_REPEAT}+1)',
setpts=N/FRAME_RATE/TB,
drawtext=fontfile=${FONT_FILE}:
text='Step %{eif\:n+${STEP_TO_REPEAT}+1\:d}':
x=w-tw-10:y=10:
fontsize=24:
fontcolor=white:
box=1:boxcolor=black@0.5" -fps_mode vfr -c:v libx264 -pix_fmt yuv420p part3_fixed_timestamps.mp4

# Create filelist.txt for concatenation
echo "file 'part1.mp4'" > filelist.txt
echo "file 'part2_final.mp4'" >> filelist.txt
echo "file 'part3_fixed_timestamps.mp4'" >> filelist.txt

# Concatenate videos
ffmpeg -f concat -safe 0 -i filelist.txt -c copy ${OUTPUT_VIDEO}

# Clean up intermediate files
rm part1.mp4 part2_final.mp4 part3_fixed_timestamps.mp4 filelist.txt

echo "Final video generated: ${OUTPUT_VIDEO}"
