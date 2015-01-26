#============================================================================
OUT_NAME    := sizeof
#EXEC_EXT    := .exe
OUT_DIR     := .
CLEAN_FILES := "$(OUT_DIR)/$(OUT_NAME).exe"
#----------------------------------------------------------------------------

# 1-st way to select source files
SRCS := sizeof.c

# 2-nd way to select source files
#SRC_DIRS := .

#----------------------------------------------------------------------------
DEFS      :=
#INC_DIRS  := ../libs/include 

OPTIM     := -Os -fomit-frame-pointer
#OPTIM     := -O0 -g
WARN      := -Wall

#ASFLAGS   := -mcpu=cortex-m3 -mthumb $(ASFLAGS)
CFLAGS    := $(WARN) -pipe $(OPTIM) $(DEFS) $(CFLAGS)
CXXFLAGS  := $(CFLAGS) $(CXXFLAGS)
LDFLAGS   := -lm $(LDFLAGS)
#----------------------------------------------------------------------------
#_AS  := @as
#_CC  := @gcc
#_CXX := @g++
#_LD  := @g++
#----------------------------------------------------------------------------
include Makefile.skel
#============================================================================
