```
title: '吴恩达ML学习笔记'
date: 2021-09-21 19:14:31
tags: [ML]
published: true
hideInList: false
feature: 
isTop: false
```

# 吴恩达ML学习笔记

## 机器学习定义

- 计算机从经验`E`中学习，解决任务`T`，进行某个性能度量`P`，通过`P`测定在`T`上的表现因经验`E`而提高。

## 机器学习分类

- Supervised learning监督学习：教会计算机做某件事情
- Unsupervised learning无监督学习：让计算机自己去学习
- Reinforcement learning强化学习
- Recommend systems推荐系统
- ···

## 监督学习

给算法一个数据集，这个数据集中包含了正确的答案，并告诉计算机什么是正确的、什么是错误的（或者说数据对应的明确标签）；算法的目的是让机器给出更多正确的答案。

### 回归问题-regression

预测连续的数值属性。

- 预测房价 

#### 单变量线性回归-Linear regression with one variable

$$
Hypothesis：			h_{\theta}(x)=\theta_0+\theta_1x
$$

$$
Parameters:		\theta_0,\theta_1
$$

$$
Cost Function:		J(\theta_0,\theta_1) = {1\over2m}\sum_{i=1}^m(h_{\theta}(x^{(i)})-y^{(i)})^2
$$

$$
Goal: minimize_{(\theta_0,\theta_1)}J(\theta_0,\theta_1)
$$

既然我们的目标是将代价函数最小化，那一个一个试参数将会非常麻烦。所以这里引入**梯度函数**，快速将代价函数`J`最小化。

#### 梯度下降算法-Gradient descent-Batch

$$
重复直至收敛：\theta_j := \theta_j - \alpha\frac{\partial }{\partial \theta_j}J(\theta_0,\theta_1){\,}{\,}(for{\,}{\,}j=0{\,}{\,}and{\,}{\,}j=1)
$$

或者换一种表达方法(同样的要进行到收敛)：
$$
\theta_0 :=\theta_0-\alpha{1\over m}\sum_{i=1}^m(h_{\theta}(x^{(i)})-y^{(i)})
$$

$$
\theta_1:=\theta_1-\alpha{1\over m}\sum_{i=1}^m(h_{\theta}(x^{(i)})-y^{(i)})·x^{(i)}
$$



其中：
$$
\alpha
$$
表示学习率，也就是梯度下降时我们迈出多大的步子。越小则说明梯度下降的速率越缓慢，越大则说明梯度下降的速率越迅速。

梯度下降是很常用的算法，它是一个一阶的最优化算法，不仅被用在线性回归上，还被用在众多的机器学习领域中。

它可以解决更一般的问题。

Have some function: 
$$
J(\theta_0,\theta_1,\theta_2,...,\theta_n)
$$
Want:
$$
min_{(\theta_0,...,\theta_n)}J(\theta_0,...,\theta_n)
$$
特点：

- 沿着不同路线下降，会有多个局部最优解，容易陷入局部最优化

#### 多元线性回归模型及梯度下降算法

$$
Hypothesis:h_\theta(x)=\theta^Tx=\theta_0x_0+\theta_1x_1+\theta_2x_2+···+\theta_nx_n
$$

$$
Parameters:\theta_0,\theta_1,···,\theta_n
$$

$$
Cost Function:J(\theta_0,\theta_1,···,\theta_n)={1\over2m}\sum_{i=1}^m(h_{\theta}(x^{(i)})-y^{(i)})^2
$$

$$
GradientDescent:Repeat\{\theta_j:=\theta_j-\alpha\frac{\partial }{\partial \theta_j}J(\theta_0,\theta_1,···,\theta_n)\}
$$

将`GradientDescent`的偏导数展开，就是：
$$
\theta_j:=\theta_j-\alpha{1\over m}\sum_{i=1}^m(h_{\theta}(x^{(i)})-y^{(i)})x^{(i)}_j
$$

#### 处理梯度下降的常用技巧

##### 特征缩放

如果一个问题有很多特征，这些特征的取值都处在一个相近的范围，那么梯度下降算法就能更快地收敛。

特征缩放的目的是：将特征的取值约束到`-1`到`+1`的范围内。

##### 归一化

进行如下替换，让特征值具有为0的平均值：
$$
x_i->(x_i-\mu_i)
$$
其中：
$$
x_i:第i个特征
$$

$$
\mu_i:第i个特征x_i的平均值
$$

然后用：
$$
x_i-\mu_i\over s_i
$$
去替换特征值：
$$
x_i
$$
其中：
$$
s_i:特征x_i的规模或者说是取值范围
$$

#### 正规方程用来最小化代价函数

除了可以使用梯度下降法求解最优代价方程的`θ`，还可以使用最小化代价函数直接求解最优`θ`。
$$
\theta = (X^TX)^{-1}X^Ty
$$
其中：
$$
X：特征矩阵
$$

$$
y:结果向量
$$

如此得到的`θ`就可以将代价函数最小化。

### 分类问题-classification

预测离散的数值属性。

- 判断肿瘤良性与否

#### Logistic regression

有$h_\theta(x)=g(\theta^Tx)$；其中$g(z)={1\over1+e^{-z}}$​​ ，又叫做"logistic function"或者“sigmoid function”。在这里

$h_\theta(x)$预测的是在参数$\theta$、特征值$x$的条件下，$y=1$​的概率。

#### Logistic regression 的 cost function

与上面回归模型不同是，在使用“logistic function”后，我们的$J(\theta)$​代价函数的图像会变成“非凸”的，会存在多个局部最小值。所以我们想找一个只有一个最值的图像。这里引入逻辑回归的代价函数：

整体的代价函数如下：
$$
J(\theta)={1\over m}\sum_{i=1}^mCost(h_\theta(x)^{(i)},y^{(i)})
$$


其中：
$$
Cost(h_\theta(x),y) = 
 \begin{cases}
 -log(h_\theta(x)),\,\,if \,\, y=1\\
-log(1-h_\theta(x)),\,\,if \,\, y=0\\
 \end{cases}
$$
这里需要注意：$y=0 \,\, or \,\, 1 \,\, always$​

所以，$Cost(h_\theta(x),y)$又可以写成：
$$
Cost(h_\theta(x),y)=-ylog(h_\theta(x))-(1-y)log(1-h_\theta(x))
$$
最终，新的代价函数$J(\theta)$就是：
$$
J(\theta)=-{1\over m}[\sum_{i=1}^my^{(i)}log(h_\theta(x^{(i)}))+(1-y^{(i)})log(1-h_\theta(x^{(i)}))]
$$
接下来，我们仍希望去最小化代价函数，得到$min_\theta J(\theta)$。

与上面回归问题相似，我们依旧使用梯度下降法：
$$
Repeat\{\\
\theta_j:=\theta_j-\alpha\frac{\partial }{\partial \theta_j}J(\theta) \\
simulataneously\,\,update\,\,all \,\,\theta_j\}
$$
**需要注意的是**：这里的公式看似与前面相同，但是我们对$h_{\theta}(x)$的定义发生了变化，所以这是完全不同的。

## 无监督学习

给算法的数据集没有明确的目的和用途，也不清楚每个数据点的意义，让计算机从中找出某种结构（比如让机器能够将数据分成若干个”簇“），

### 聚类算法-clustering algorithm

告诉算法，这有一堆数据，不知道这些数据是什么、不知道谁是什么类型、甚至不知道有哪些类型（当然也无法告诉机器什么是正确答案），让机器自动找出这些数据的结构并按照得到的结构类型将这些数据个体分成”簇“

- Google news
- genes自动分类
- 自动管理计算机集群
- 社会网络分析
- 星际数据分析



## 常见问题

### 过拟合问题

过拟合问题通常在变量过多时出现，在训练时的假设可以很好的拟合训练集，代价函数实际上很可能接近于甚至等于0，这样一来模型会千方百计的去拟合训练集，最终导致**模型无法泛化到新的样本中**。

#### 如何解决过拟合

1. 减少特征的数量

   - 人工的去决定哪些特征变量是重要的。
   - 使用模型选择算法，让算法去决定保留哪些特征变量。

   这种方法的缺点是：会舍弃一些信息，尽管这些信息是有用的。

2. 正则化Regularization

   - 保持所有的特征变量，但是减少量级或者参数$\theta$的大小。
   - 这样就保留了所有特征变量，因为每一个变量都会对预测的模型产生或大或小的影响。

### 欠拟合问题

与过拟合相反。

## 符号定义

$$
m：训练样本的数量
$$


$$
x's：输入变量，或者说是特征
$$

$$
y's：输出变量，也就是要预测的目标变量
$$

$$
(x,y)：表示一个训练样本
$$

$$
(x^{(i)},y^{(i)})：表示第i个训练样本
$$

$$
h：假设函数(hypothesis)，接收x，尝试输出y
$$

$$
假设函数h_{\theta}(x)=\theta_0+\theta_1x，那么\theta_0和\theta_1叫做模型参数
$$

























