package bstream

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	inputBytes = []byte{0, 1, 36, 128, 164, 255}
)

func Test_bigEndian_ReadInt16Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			var (
				inp  = [2]byte{byte1, byte2}
				want = int16(binary.BigEndian.Uint16(inp[:]))
			)
			assert.Equalf(t, want, bigEndian{}.ReadInt16Array(inp), "ReadInt16Array(%#v)", inp)
		}
	}
}

func Test_bigEndian_ReadInt32Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					var (
						inp  = [4]byte{byte1, byte2, byte3, byte4}
						want = int32(binary.BigEndian.Uint32(inp[:]))
					)
					assert.Equalf(t, want, bigEndian{}.ReadInt32Array(inp), "ReadInt32Array(%#v)", inp)
				}
			}
		}
	}
}

func Test_bigEndian_ReadInt64Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					for _, byte5 := range inputBytes {
						for _, byte6 := range inputBytes {
							for _, byte7 := range inputBytes {
								for _, byte8 := range inputBytes {
									var (
										inp  = [8]byte{byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8}
										want = int64(binary.BigEndian.Uint64(inp[:]))
									)
									assert.Equalf(t, want, bigEndian{}.ReadInt64Array(inp), "ReadInt64Array(%#v)", inp)
								}
							}
						}
					}
				}
			}
		}
	}
}

func Test_bigEndian_ReadUint16Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			var (
				inp  = [2]byte{byte1, byte2}
				want = binary.BigEndian.Uint16(inp[:])
			)
			assert.Equalf(t, want, bigEndian{}.ReadUint16Array(inp), "ReadUint16Array(%#v)", inp)
		}
	}
}

func Test_bigEndian_ReadUint32Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					var (
						inp  = [4]byte{byte1, byte2, byte3, byte4}
						want = binary.BigEndian.Uint32(inp[:])
					)
					assert.Equalf(t, want, bigEndian{}.ReadUint32Array(inp), "ReadUint32Array(%#v)", inp)
				}
			}
		}
	}
}

func Test_bigEndian_ReadUint64Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					for _, byte5 := range inputBytes {
						for _, byte6 := range inputBytes {
							for _, byte7 := range inputBytes {
								for _, byte8 := range inputBytes {
									var (
										inp  = [8]byte{byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8}
										want = binary.BigEndian.Uint64(inp[:])
									)
									assert.Equalf(t, want, bigEndian{}.ReadUint64Array(inp), "ReadUint64Array(%#v)",
										inp)
								}
							}
						}
					}
				}
			}
		}
	}
}

func Test_bigEndian_WriteInt16Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			var (
				inp  = [2]byte{byte1, byte2}
				want = make([]byte, 2)
				arg  = int16(binary.BigEndian.Uint16(inp[:]))
			)
			binary.BigEndian.PutUint16(want, uint16(arg))
			assert.EqualValues(t, *(*[2]byte)(want), bigEndian{}.WriteInt16Array(arg), "ReadInt16Array(%#v)", inp)
		}
	}
}

func Test_bigEndian_WriteInt32Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					var (
						inp  = [4]byte{byte1, byte2, byte3, byte4}
						want = make([]byte, 4)
						arg  = int32(binary.BigEndian.Uint32(inp[:]))
					)
					binary.BigEndian.PutUint32(want, uint32(arg))
					assert.EqualValues(t, *(*[4]byte)(want), bigEndian{}.WriteInt32Array(arg), "ReadInt32Array(%#v)",
						inp)
				}
			}
		}
	}
}

func Test_bigEndian_WriteInt64Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					for _, byte5 := range inputBytes {
						for _, byte6 := range inputBytes {
							for _, byte7 := range inputBytes {
								for _, byte8 := range inputBytes {
									var (
										inp  = [8]byte{byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8}
										want = make([]byte, 8)
										arg  = int64(binary.BigEndian.Uint64(inp[:]))
									)
									binary.BigEndian.PutUint64(want, uint64(arg))
									assert.EqualValues(t, *(*[8]byte)(want), bigEndian{}.WriteInt64Array(arg),
										"ReadInt64Array(%#v)", inp)
								}
							}
						}
					}
				}
			}
		}
	}
}

func Test_bigEndian_WriteUint16Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			var (
				inp  = [2]byte{byte1, byte2}
				want = make([]byte, 2)
				arg  = binary.BigEndian.Uint16(inp[:])
			)
			binary.BigEndian.PutUint16(want, arg)
			assert.EqualValues(t, *(*[2]byte)(want), bigEndian{}.WriteUint16Array(arg), "ReadUint16Array(%#v)", inp)
		}
	}
}

func Test_bigEndian_WriteUint32Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					var (
						inp  = [4]byte{byte1, byte2, byte3, byte4}
						want = make([]byte, 4)
						arg  = binary.BigEndian.Uint32(inp[:])
					)
					binary.BigEndian.PutUint32(want, arg)
					assert.EqualValues(t, *(*[4]byte)(want), bigEndian{}.WriteUint32Array(arg), "ReadUint32Array(%#v)",
						inp)
				}
			}
		}
	}
}

func Test_bigEndian_WriteUint64Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					for _, byte5 := range inputBytes {
						for _, byte6 := range inputBytes {
							for _, byte7 := range inputBytes {
								for _, byte8 := range inputBytes {
									var (
										inp  = [8]byte{byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8}
										want = make([]byte, 8)
										arg  = binary.BigEndian.Uint64(inp[:])
									)
									binary.BigEndian.PutUint64(want, arg)
									assert.EqualValues(t, *(*[8]byte)(want), bigEndian{}.WriteUint64Array(arg),
										"ReadUint64Array(%#v)", inp)
								}
							}
						}
					}
				}
			}
		}
	}
}

func Test_littleEndian_ReadInt16Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			var (
				inp  = [2]byte{byte1, byte2}
				want = int16(binary.LittleEndian.Uint16(inp[:]))
			)
			assert.Equalf(t, want, littleEndian{}.ReadInt16Array(inp), "ReadInt16Array(%#v)", inp)
		}
	}
}

func Test_littleEndian_ReadInt32Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					var (
						inp  = [4]byte{byte1, byte2, byte3, byte4}
						want = int32(binary.LittleEndian.Uint32(inp[:]))
					)
					assert.Equalf(t, want, littleEndian{}.ReadInt32Array(inp), "ReadInt32Array(%#v)", inp)
				}
			}
		}
	}
}

func Test_littleEndian_ReadInt64Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					for _, byte5 := range inputBytes {
						for _, byte6 := range inputBytes {
							for _, byte7 := range inputBytes {
								for _, byte8 := range inputBytes {
									var (
										inp  = [8]byte{byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8}
										want = int64(binary.LittleEndian.Uint64(inp[:]))
									)
									assert.Equalf(t, want, littleEndian{}.ReadInt64Array(inp), "ReadInt64Array(%#v)",
										inp)
								}
							}
						}
					}
				}
			}
		}
	}
}

func Test_littleEndian_ReadUint16Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			var (
				inp  = [2]byte{byte1, byte2}
				want = binary.LittleEndian.Uint16(inp[:])
			)
			assert.Equalf(t, want, littleEndian{}.ReadUint16Array(inp), "ReadUint16Array(%#v)", inp)
		}
	}
}

func Test_littleEndian_ReadUint32Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					var (
						inp  = [4]byte{byte1, byte2, byte3, byte4}
						want = binary.LittleEndian.Uint32(inp[:])
					)
					assert.Equalf(t, want, littleEndian{}.ReadUint32Array(inp), "ReadUint32Array(%#v)", inp)
				}
			}
		}
	}
}

func Test_littleEndian_ReadUint64Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					for _, byte5 := range inputBytes {
						for _, byte6 := range inputBytes {
							for _, byte7 := range inputBytes {
								for _, byte8 := range inputBytes {
									var (
										inp  = [8]byte{byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8}
										want = binary.LittleEndian.Uint64(inp[:])
									)
									assert.Equalf(t, want, littleEndian{}.ReadUint64Array(inp), "ReadUint64Array(%#v)",
										inp)
								}
							}
						}
					}
				}
			}
		}
	}
}

func Test_littleEndian_WriteInt16Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			var (
				inp  = [2]byte{byte1, byte2}
				want = make([]byte, 2)
				arg  = int16(binary.LittleEndian.Uint16(inp[:]))
			)
			binary.LittleEndian.PutUint16(want, uint16(arg))
			assert.EqualValues(t, *(*[2]byte)(want), littleEndian{}.WriteInt16Array(arg), "ReadInt16Array(%#v)", inp)
		}
	}
}

func Test_littleEndian_WriteInt32Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					var (
						inp  = [4]byte{byte1, byte2, byte3, byte4}
						want = make([]byte, 4)
						arg  = int32(binary.LittleEndian.Uint32(inp[:]))
					)
					binary.LittleEndian.PutUint32(want, uint32(arg))
					assert.EqualValues(t, *(*[4]byte)(want), littleEndian{}.WriteInt32Array(arg), "ReadInt32Array(%#v)",
						inp)
				}
			}
		}
	}
}

func Test_littleEndian_WriteInt64Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					for _, byte5 := range inputBytes {
						for _, byte6 := range inputBytes {
							for _, byte7 := range inputBytes {
								for _, byte8 := range inputBytes {
									var (
										inp  = [8]byte{byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8}
										want = make([]byte, 8)
										arg  = int64(binary.LittleEndian.Uint64(inp[:]))
									)
									binary.LittleEndian.PutUint64(want, uint64(arg))
									assert.EqualValues(t, *(*[8]byte)(want), littleEndian{}.WriteInt64Array(arg),
										"ReadInt64Array(%#v)", inp)
								}
							}
						}
					}
				}
			}
		}
	}
}

func Test_littleEndian_WriteUint16Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			var (
				inp  = [2]byte{byte1, byte2}
				want = make([]byte, 2)
				arg  = binary.LittleEndian.Uint16(inp[:])
			)
			binary.LittleEndian.PutUint16(want, arg)
			assert.EqualValues(t, *(*[2]byte)(want), littleEndian{}.WriteUint16Array(arg), "ReadUint16Array(%#v)", inp)
		}
	}
}

func Test_littleEndian_WriteUint32Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					var (
						inp  = [4]byte{byte1, byte2, byte3, byte4}
						want = make([]byte, 4)
						arg  = binary.LittleEndian.Uint32(inp[:])
					)
					binary.LittleEndian.PutUint32(want, arg)
					assert.EqualValues(t, *(*[4]byte)(want), littleEndian{}.WriteUint32Array(arg),
						"ReadUint32Array(%#v)", inp)
				}
			}
		}
	}
}

func Test_littleEndian_WriteUint64Array(t *testing.T) {
	for _, byte1 := range inputBytes {
		for _, byte2 := range inputBytes {
			for _, byte3 := range inputBytes {
				for _, byte4 := range inputBytes {
					for _, byte5 := range inputBytes {
						for _, byte6 := range inputBytes {
							for _, byte7 := range inputBytes {
								for _, byte8 := range inputBytes {
									var (
										inp  = [8]byte{byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8}
										want = make([]byte, 8)
										arg  = binary.LittleEndian.Uint64(inp[:])
									)
									binary.LittleEndian.PutUint64(want, arg)
									assert.EqualValues(t, *(*[8]byte)(want), littleEndian{}.WriteUint64Array(arg),
										"ReadUint64Array(%#v)", inp)
								}
							}
						}
					}
				}
			}
		}
	}
}
