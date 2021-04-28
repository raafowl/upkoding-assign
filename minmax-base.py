class n:
  LOWER_BASE = 0
  UPPER_BASE = 9

  def __init__(self, n):
    self.n = n

  def minBase(self):
    minn = self.UPPER_BASE
    for char in str(self.n):
      if minn > int(char):
        minn = int(char)
    return minn

  def maxBase(self):
    maxn = self.LOWER_BASE
    for char in str(self.n):
      if maxn < int(char):
        maxn = int(char)
    return maxn

myNumber = n(23)
print('Max: {}'.format(myNumber.maxBase()))
print('Min: {}'.format(myNumber.minBase()))
