# coding=utf-8
# 处理数据，生成图像，显示图像并保存图像
import os
import sys

import numpy as np
import matplotlib.pyplot as plt
from mpl_toolkits.axes_grid1 import AxesGrid


def draw_grid(fig, data_array, x_min, x_max, y_min, y_max):
    """
    A grid of 2x2 images with a single colorbar
    """
    grid = AxesGrid(fig, 111,  # similar to subplot(142)
                    nrows_ncols=(1, 1),
                    axes_pad=0.2,
                    share_all=True,
                    label_mode="L",
                    cbar_location="right",
                    cbar_mode="single",
                    )
    Z = data_array
    extent = (x_min, x_max, y_min, y_max)

    im = grid[0].imshow(Z, extent=extent, interpolation="nearest")
    grid[0].set_aspect(2.0, adjustable='box')

    grid.cbar_axes[0].colorbar(im)

    for cax in grid.cbar_axes:
        cax.toggle_label(True)

    # This affects all axes as share_all = True.
    grid.axes_llc.set_xticks(range(x_min, x_max, 10))  # Set x-axes sequence
    grid.axes_llc.set_yticks(range(y_min, y_max, 10))


def is_valid_data(data):
    """
    Tells the input data is valid or not.
    """
    if isinstance(data, str):
        data = float(data)

    return data < 19.0


# Check arguments
if len(sys.argv) < 3:
    usage = '''
    Usage 1: python D:\py\draw.py input_file output_file [figure_file]
    input_file: Path of the input text file
    output_file: Path of the output text file. File will be OVERWRITTEN if it already exists
    figure_file: Path of the figure file(should be *.png). This is optional and will be result.png if not given.
                 File will be OVERWRITTEN if it already exists
    ===== E.T. for W.W., July 3rd, 2017 ======
    '''
    print "Invalid Argument"
    print usage
    sys.exit(1)

save_figure_file_name = 'result.png'
if len(sys.argv) >= 4:
    save_figure_file_name = sys.argv[3]

data_metrics = []
with open(sys.argv[1], 'r') as ifile:
    with open(sys.argv[2], 'w') as ofile:
        row_array = []
        lines = ifile.readlines()
        for idx in range(0, len(lines)):
            try:
                line = lines[idx]
                if len(line) <= 2:
                    print '#', idx, 'Skip possible blank line:', line
                    continue
                line = line[:-1]  # Remove \n
                parts = line.split()
                if len(parts) != 2:
                    print '#', idx, 'Possible bad line:', line
                
                pos, data = line.split() # Split by space

                if len(pos) != 4:
                    print '#', idx, 'Possible bad line:', line
                
                row = int(pos[:2])
                col = int(pos[2:4])

                if not is_valid_data(data):
                    data = '0'

                row_array.append(float(data))
                ofile.write(data + "\t")
                if col == 99:
                    ofile.write("\n")
                    data_metrics.append(row_array)
                    row_array = []
                
            except Exception as e:
                print '#', idx, 'Exception:', str(e)

if len(data_metrics) != 50:
    print "Rows not reaching 50, exit"
    sys.exit(1)

np_array = np.array(data_metrics, dtype=float)
F = plt.figure(1, (8, 8))  # Size of the figure
draw_grid(F, np_array, 0, 100, 0, 50)
plt.draw()
plt.savefig('result.png', dpi=72)
plt.show()
