

def calc_sum(num)
    sum = 0
    while num > 0
        remain = num % 10
        sum += remain**2
        num /= 10
    end
    sum
end

def is_happy(num)
    h = {}
    while (num = calc_sum(num)) != 1
        if h.has_key? num
            return false
        else
            h[num] = true
        end
    end
    true
end

p is_happy(19) # should be true
p is_happy(2345) # should be false