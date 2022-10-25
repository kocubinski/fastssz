// Code generated by fastssz. DO NOT EDIT.
// Hash: af56ac74ccaeb89b8d2d802f56cbfab0248f088a80a7879c0e2f0fb8ca88f22c
// Version: 0.1.3-dev
package testcases

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Case7 object
func (c *Case7) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the Case7 object to a target array
func (c *Case7) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'BlobKzgs'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(c.BlobKzgs) * 48

	// Field (0) 'BlobKzgs'
	if size := len(c.BlobKzgs); size > 16 {
		err = ssz.ErrListTooBigFn("Case7.BlobKzgs", size, 16)
		return
	}
	for ii := 0; ii < len(c.BlobKzgs); ii++ {
		if size := len(c.BlobKzgs[ii]); size != 48 {
			err = ssz.ErrBytesLengthFn("Case7.BlobKzgs[ii]", size, 48)
			return
		}
		dst = append(dst, c.BlobKzgs[ii]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Case7 object
func (c *Case7) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'BlobKzgs'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 4 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (0) 'BlobKzgs'
	{
		buf = tail[o0:]
		num, err := ssz.DivideInt2(len(buf), 48, 16)
		if err != nil {
			return err
		}
		c.BlobKzgs = make([][]byte, num)
		for ii := 0; ii < num; ii++ {
			if cap(c.BlobKzgs[ii]) == 0 {
				c.BlobKzgs[ii] = make([]byte, 0, len(buf[ii*48:(ii+1)*48]))
			}
			c.BlobKzgs[ii] = append(c.BlobKzgs[ii], buf[ii*48:(ii+1)*48]...)
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Case7 object
func (c *Case7) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'BlobKzgs'
	size += len(c.BlobKzgs) * 48

	return
}

// HashTreeRoot ssz hashes the Case7 object
func (c *Case7) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the Case7 object with a hasher
func (c *Case7) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'BlobKzgs'
	{
		if size := len(c.BlobKzgs); size > 16 {
			err = ssz.ErrListTooBigFn("Case7.BlobKzgs", size, 16)
			return
		}
		subIndx := hh.Index()
		for _, i := range c.BlobKzgs {
			if len(i) != 48 {
				err = ssz.ErrBytesLength
				return
			}
			hh.PutBytes(i)
		}
		numItems := uint64(len(c.BlobKzgs))
		hh.MerkleizeWithMixin(subIndx, numItems, 16)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Case7 object
func (c *Case7) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}
