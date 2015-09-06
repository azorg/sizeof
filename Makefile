#============================================================================
OUT_NAME    := sizeof
#EXEC_EXT   := .exe
OUT_DIR     := .
#CLEAN_DIR  := $(OUT_DIR)/tmp
CLEAN_FILES := "$(OUT_DIR)/$(OUT_NAME).exe"
#----------------------------------------------------------------------------
# 1-st way to select source files
SRCS := sizeof.c
HDRS :=

# 2-nd way to select source files
#SRC_DIRS := .
#HDR_DIRS :=
#----------------------------------------------------------------------------
#INC_DIRS  := ../libs/include 
#DEFS      := -DSIZEOF_DEBUG
OPTIM      := -Os -fomit-frame-pointer
#OPTIM     := -O0 -g
WARN       := -Wall

#ASFLAGS   := -mcpu=cortex-m3 -mthumb $(ASFLAGS)
CFLAGS     := $(WARN) $(OPTIM) $(DEFS) $(CFLAGS) -pipe
CXXFLAGS   := $(CXXFLAGS) $(CFLAGS)
LDFLAGS    := -lm $(LDFLAGS)
#----------------------------------------------------------------------------
#_AS  := @as
#_CC  := @gcc
#_CXX := @g++
#_LD  := @gcc

#_CC  := @clang
#_CXX := @clang++
#_LD  := @clang
#----------------------------------------------------------------------------
include Makefile.skel
#============================================================================
