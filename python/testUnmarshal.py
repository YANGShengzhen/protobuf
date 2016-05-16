import proto
try:
    f = open('../proto/data/1', 'r')
    data = f.read()
    message = proto.ResponseWrapper.FromString(data)
    print message
    print len(data)
finally:
    if f:
        f.close()
